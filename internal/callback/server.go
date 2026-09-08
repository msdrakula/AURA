// Package callback is a webhook receiver that logs incoming HTTP requests
// grouped by a request ID extracted from the first path segment.
package callback

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Interaction is one captured inbound request.
type Interaction struct {
	Timestamp  time.Time         `json:"timestamp"`
	RemoteAddr string            `json:"remote_addr"`
	Method     string            `json:"method"`
	URL        string            `json:"url"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

// Server listens for inbound HTTP requests and stores them keyed by request ID.
type Server struct {
	port         int
	interactions *sync.Map
	logger       *zap.Logger
	httpServer   *http.Server
	listener     net.Listener
	mu           sync.Mutex
}

// NewServer creates a callback server bound to port (0 = random free port).
func NewServer(port int, logger *zap.Logger) *Server {
	if logger == nil {
		logger = zap.NewNop()
	}
	return &Server{
		port:         port,
		interactions: &sync.Map{},
		logger:       logger,
	}
}

// Addr returns the bound address, or empty before Start.
func (s *Server) Addr() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.listener == nil {
		return ""
	}
	return s.listener.Addr().String()
}

// Start binds the listener and serves until ctx is cancelled or the server fails.
func (s *Server) Start(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}
	s.mu.Lock()
	if s.listener != nil {
		s.mu.Unlock()
		return fmt.Errorf("callback: already started")
	}
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", s.port))
	if err != nil {
		s.mu.Unlock()
		return fmt.Errorf("callback: listen on port %d: %w", s.port, err)
	}
	s.listener = ln
	s.httpServer = &http.Server{
		Handler:           http.HandlerFunc(s.handle),
		ReadHeaderTimeout: 10 * time.Second,
		BaseContext:       func(net.Listener) context.Context { return ctx },
	}
	s.mu.Unlock()

	s.logger.Info("callback listening", zap.String("addr", s.Addr()))

	errCh := make(chan error, 1)
	go func() {
		errCh <- s.httpServer.Serve(ln)
	}()

	select {
	case <-ctx.Done():
		sh, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.httpServer.Shutdown(sh); err != nil {
			s.logger.Warn("callback shutdown", zap.Error(err))
		}
		<-errCh
		s.mu.Lock()
		s.listener = nil
		s.httpServer = nil
		s.mu.Unlock()
		return ctx.Err()
	case err := <-errCh:
		s.mu.Lock()
		s.listener = nil
		s.httpServer = nil
		s.mu.Unlock()
		if err == nil || err == http.ErrServerClosed {
			return nil
		}
		return fmt.Errorf("callback: serve: %w", err)
	}
}

func (s *Server) handle(w http.ResponseWriter, r *http.Request) {
	requestID := extractRequestID(r.URL.Path)
	body, _ := io.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()

	headers := make(map[string]string, len(r.Header))
	for k, v := range r.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	interaction := Interaction{
		Timestamp:  time.Now().UTC(),
		RemoteAddr: r.RemoteAddr,
		Method:     r.Method,
		URL:        r.URL.String(),
		Headers:    headers,
		Body:       string(body),
	}

	actual, _ := s.interactions.LoadOrStore(requestID, []Interaction{})
	updated := append(actual.([]Interaction), interaction)
	s.interactions.Store(requestID, updated)

	s.logger.Debug("callback hit",
		zap.String("request_id", requestID),
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
		zap.Int("body_bytes", len(body)),
	)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

// extractRequestID takes the first non-empty path segment, or generates a UUID
// for the root path.
func extractRequestID(path string) string {
	trimmed := strings.TrimPrefix(path, "/")
	if trimmed == "" {
		return uuid.NewString()
	}
	if i := strings.IndexByte(trimmed, '/'); i >= 0 {
		trimmed = trimmed[:i]
	}
	return trimmed
}

// GetInteractions returns all stored interactions for requestID, or nil if none.
func (s *Server) GetInteractions(requestID string) []Interaction {
	if requestID == "" {
		return nil
	}
	v, ok := s.interactions.Load(requestID)
	if !ok {
		return nil
	}
	return v.([]Interaction)
}
