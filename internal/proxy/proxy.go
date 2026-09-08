// Package proxy implements a local HTTP/HTTPS intercepting proxy.
package proxy

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"

	"meb/internal/analyzer"
	"meb/internal/certs"
	"meb/internal/codec"
	"meb/internal/httpio"
)

const (
	defaultIdleTimeout = 60 * time.Second
	defaultDialTimeout = 30 * time.Second
)

// Options wires runtime intercept state, optional history, and certificates.
type Options struct {
	Runtime  Runtime
	History  TransactionWriter
	Findings FindingWriter
	Analyzer *analyzer.Engine
	Certs    *certs.Authority
	Log      *zap.Logger
}

// Server is a local HTTP/1.1 proxy with optional TLS interception
// for connections that use CONNECT.
type Server struct {
	rt       Runtime
	history  TransactionWriter
	findings FindingWriter
	analyzer *analyzer.Engine
	certs    *certs.Authority
	log      *zap.Logger
	idle     time.Duration
	dial     time.Duration

	mu     sync.Mutex
	ln     net.Listener
	addr   string
	cancel context.CancelFunc
	conns  map[net.Conn]struct{}
}

// New constructs a proxy from Options. Log may be nil.
func New(opts Options) *Server {
	log := opts.Log
	if log == nil {
		log = zap.NewNop()
	}
	eng := opts.Analyzer
	if eng == nil {
		eng = analyzer.DefaultEngine()
	}
	return &Server{
		rt:       opts.Runtime,
		history:  opts.History,
		findings: opts.Findings,
		analyzer: eng,
		certs:    opts.Certs,
		log:      log,
		idle:     defaultIdleTimeout,
		dial:     defaultDialTimeout,
		conns:    map[net.Conn]struct{}{},
	}
}

// Addr returns the bound listen address, or empty if the proxy is stopped.
func (s *Server) Addr() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.addr
}

// Start begins accepting connections. ctx cancellation stops the listener.
// The HTTP API should pass a process-scoped context, not the request context.
func (s *Server) Start(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}
	s.mu.Lock()
	if s.ln != nil {
		s.mu.Unlock()
		return nil
	}
	if s.rt == nil {
		s.mu.Unlock()
		return fmt.Errorf("proxy: runtime store is not configured")
	}
	if s.certs == nil {
		s.mu.Unlock()
		return fmt.Errorf("proxy: certificates are not configured")
	}
	if err := s.certs.Ensure(); err != nil {
		s.mu.Unlock()
		return err
	}
	host, port := s.rt.ListenAddr()
	addr := net.JoinHostPort(host, strconv.Itoa(port))
	lc := net.ListenConfig{}
	ln, err := lc.Listen(ctx, "tcp", addr)
	if err != nil {
		s.mu.Unlock()
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}
	runCtx, cancel := context.WithCancel(ctx)
	s.ln = ln
	s.addr = ln.Addr().String()
	s.cancel = cancel
	s.mu.Unlock()

	s.rt.SetRunning(true)
	s.rt.Log("info", "proxy listening on "+s.addr)
	s.log.Info("proxy listening", zap.String("addr", s.addr))
	go s.accept(runCtx, ln)
	return nil
}

// Stop closes the listener and in-flight client connections.
func (s *Server) Stop(ctx context.Context) error {
	_ = ctx
	s.mu.Lock()
	ln := s.ln
	cancel := s.cancel
	s.ln = nil
	s.addr = ""
	s.cancel = nil
	conns := make([]net.Conn, 0, len(s.conns))
	for c := range s.conns {
		conns = append(conns, c)
	}
	s.conns = map[net.Conn]struct{}{}
	s.mu.Unlock()

	if s.rt != nil {
		s.rt.SetRunning(false)
		s.rt.ForwardAll()
	}
	if cancel != nil {
		cancel()
	}
	if ln != nil {
		if err := ln.Close(); err != nil {
			s.log.Warn("listener close", zap.Error(err))
		}
	}
	for _, c := range conns {
		_ = c.Close()
	}
	if s.rt != nil {
		s.rt.Log("info", "proxy stopped")
	}
	s.log.Info("proxy stopped")
	return nil
}

func (s *Server) track(c net.Conn) {
	s.mu.Lock()
	s.conns[c] = struct{}{}
	s.mu.Unlock()
}

func (s *Server) untrack(c net.Conn) {
	s.mu.Lock()
	delete(s.conns, c)
	s.mu.Unlock()
}

func (s *Server) accept(ctx context.Context, ln net.Listener) {
	go func() {
		<-ctx.Done()
		_ = ln.Close()
	}()
	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				s.log.Warn("accept", zap.Error(err))
				return
			}
		}
		s.mu.Lock()
		alive := s.ln == ln
		s.mu.Unlock()
		if !alive {
			_ = conn.Close()
			return
		}
		s.track(conn)
		go func(c net.Conn) {
			defer s.untrack(c)
			s.handle(ctx, c)
		}(conn)
	}
}

func (s *Server) handle(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	client := conn.RemoteAddr().String()
	if err := conn.SetDeadline(time.Now().Add(s.idle)); err != nil {
		s.log.Debug("set deadline", zap.Error(err))
	}
	br := bufio.NewReader(conn)
	req, err := httpio.ReadRequest(br, "http")
	if err != nil || req == nil {
		return
	}
	defer func() {
		if rec := recover(); rec != nil {
			if s.rt != nil {
				s.rt.Log("error", fmt.Sprintf("%s: %v", client, rec))
				s.rt.IncErrors()
			}
			s.log.Error("proxy panic", zap.String("client", client), zap.Any("recover", rec))
		}
	}()
	if req.Method == "CONNECT" {
		s.handleHTTPS(ctx, req, conn, br, client)
		return
	}
	s.handleHTTP(ctx, req, conn, br, client)
}

func (s *Server) handleHTTPS(ctx context.Context, connect *httpio.Request, conn net.Conn, br *bufio.Reader, client string) {
	host, port := httpio.ParseAuthority(connect.Path, 443)
	if _, err := conn.Write([]byte("HTTP/1.1 200 Connection Established\r\nProxy-Agent: MEB\r\n\r\n")); err != nil {
		return
	}
	leaf, err := s.certs.Leaf(host)
	if err != nil {
		if s.rt != nil {
			s.rt.Log("error", "TLS cert "+host+": "+err.Error())
		}
		s.log.Error("leaf cert", zap.String("host", host), zap.Error(err))
		return
	}
	cfg := &tls.Config{
		Certificates: []tls.Certificate{*leaf},
		MinVersion:   tls.VersionTLS12,
		NextProtos:   []string{"http/1.1"},
	}
	tlsConn := tls.Server(readerConn{Conn: conn, r: br}, cfg)
	if err := tlsConn.HandshakeContext(ctx); err != nil {
		if s.rt != nil {
			s.rt.Log("error", "TLS handshake "+host+": "+err.Error())
		}
		s.log.Error("tls handshake", zap.String("host", host), zap.Error(err))
		return
	}
	defer tlsConn.Close()
	_ = tlsConn.SetDeadline(time.Now().Add(s.idle))
	r := bufio.NewReader(tlsConn)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		req, err := httpio.ReadRequest(r, "https")
		if err != nil || req == nil {
			return
		}
		_ = tlsConn.SetDeadline(time.Now().Add(s.idle))
		req.Scheme = "https"
		if req.Host == "" {
			req.Host = host
			req.Port = port
		} else if req.Port == 0 {
			req.Port = port
		}
		if !s.exchange(ctx, req, tlsConn, r, client) {
			return
		}
	}
}

func (s *Server) handleHTTP(ctx context.Context, req *httpio.Request, conn net.Conn, br *bufio.Reader, client string) {
	req.Scheme = "http"
	keep := s.exchange(ctx, req, conn, br, client)
	for keep {
		select {
		case <-ctx.Done():
			return
		default:
		}
		_ = conn.SetDeadline(time.Now().Add(s.idle))
		nxt, err := httpio.ReadRequest(br, "http")
		if err != nil || nxt == nil {
			return
		}
		keep = s.exchange(ctx, nxt, conn, br, client)
	}
}

func (s *Server) exchange(ctx context.Context, req *httpio.Request, clientConn net.Conn, clientBr *bufio.Reader, client string) bool {
	_ = client
	id := newID()
	started := time.Now()
	comment := ""
	var resp *httpio.Response
	defer func() {
		s.finish(id, req, resp, started, comment)
	}()

	req.DropHeaders("proxy-connection", "proxy-authorization")
	if ae := req.Header("Accept-Encoding"); ae != "" {
		req.SetHeader("Accept-Encoding", "gzip, deflate, identity")
	}
	action, edited := s.rt.InterceptRequest(id, req)
	if action != "forward" {
		comment = "dropped"
		_ = clientConn.Close()
		return false
	}
	if edited != "" {
		applyRequestRaw(req, edited)
	}

	up, err := s.openUpstream(ctx, req)
	if err != nil {
		comment = err.Error()
		s.rt.IncErrors()
		s.log.Warn("upstream", zap.String("host", req.Host), zap.Error(err))
		_, _ = clientConn.Write(errorPage(502, comment))
		return false
	}
	defer up.Close()

	if _, err := up.Write(req.Raw()); err != nil {
		comment = fmt.Errorf("failed to write upstream: %w", err).Error()
		s.rt.IncErrors()
		return false
	}
	upBr := bufio.NewReader(up)
	if httpio.LooksLikeUpgrade(req.Headers) {
		got, err := httpio.ReadResponse(upBr, req.Method)
		if err != nil || got == nil {
			comment = "empty upstream response"
			return false
		}
		resp = got
		if _, err := clientConn.Write(resp.Raw()); err != nil {
			comment = err.Error()
			return false
		}
		if resp.Status == 101 {
			codec.Pipe(clientConn, up, leftover(clientBr), leftover(upBr))
			return false
		}
		return keepAlive(req, comment)
	}
	got, err := httpio.ReadResponse(upBr, req.Method)
	if err != nil || got == nil {
		comment = "empty upstream response"
		if err != nil {
			comment = fmt.Errorf("failed to read upstream response: %w", err).Error()
		}
		return false
	}
	codec.UnwrapBody(got)
	resp = got
	decision, editedResp := s.rt.InterceptResponse(id, req, resp)
	if decision != "forward" {
		comment = "response dropped"
		_ = up.Close()
		_ = clientConn.Close()
		return false
	}
	if editedResp != "" {
		applyResponseRaw(resp, editedResp)
	}
	if _, err := clientConn.Write(resp.Raw()); err != nil {
		comment = fmt.Errorf("failed to write client response: %w", err).Error()
		s.rt.IncErrors()
	}
	return keepAlive(req, comment)
}

func keepAlive(req *httpio.Request, comment string) bool {
	connHdr := strings.ToLower(req.Header("Connection"))
	return connHdr != "close" && comment == ""
}

func leftover(br *bufio.Reader) io.Reader {
	n := br.Buffered()
	if n <= 0 {
		return nil
	}
	buf := make([]byte, n)
	_, _ = io.ReadFull(br, buf)
	return bytes.NewReader(buf)
}

func (s *Server) openUpstream(ctx context.Context, req *httpio.Request) (net.Conn, error) {
	port := req.Port
	if port == 0 {
		if req.Scheme == "https" {
			port = 443
		} else {
			port = 80
		}
	}
	addr := net.JoinHostPort(req.Host, strconv.Itoa(port))
	d := net.Dialer{Timeout: s.dial}
	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial %s: %w", addr, err)
	}
	if err := conn.SetDeadline(time.Now().Add(s.idle)); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("failed to set deadline %s: %w", addr, err)
	}
	if req.Scheme != "https" {
		return conn, nil
	}
	verify := true
	if s.rt != nil {
		verify = s.rt.VerifyUpstream()
	}
	cfg := &tls.Config{
		ServerName:         req.Host,
		InsecureSkipVerify: !verify,
		NextProtos:         []string{"http/1.1"},
		MinVersion:         tls.VersionTLS12,
	}
	tc := tls.Client(conn, cfg)
	if err := tc.HandshakeContext(ctx); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("tls handshake %s: %w", req.Host, err)
	}
	return tc, nil
}

type readerConn struct {
	net.Conn
	r io.Reader
}

func (c readerConn) Read(p []byte) (int, error) {
	return c.r.Read(p)
}
