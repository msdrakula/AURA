package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"meb/internal/analyzer"
	"meb/internal/api"
	"meb/internal/batch"
	"meb/internal/callback"
	"meb/internal/certs"
	"meb/internal/debuglog"
	"meb/internal/desktop"
	"meb/internal/intel"
	"meb/internal/proxy"
	"meb/internal/storage"
	"meb/internal/store"
	"meb/internal/wordlist"
)

func findRoot() string {
	var cands []string
	if v := strings.TrimSpace(os.Getenv("AURA_ROOT")); v != "" {
		cands = append(cands, v)
	}
	if v := strings.TrimSpace(os.Getenv("MEB_ROOT")); v != "" {
		cands = append(cands, v)
	}
	if exe, err := os.Executable(); err == nil {
		cands = append(cands, filepath.Dir(exe))
		if real, err := filepath.EvalSymlinks(exe); err == nil {
			cands = append(cands, filepath.Dir(real))
		}
	}
	if wd, err := os.Getwd(); err == nil {
		cands = append(cands, wd)
	}
	for _, c := range cands {
		abs, err := filepath.Abs(c)
		if err != nil {
			continue
		}
		if st, err := os.Stat(filepath.Join(abs, "web", "index.html")); err == nil && !st.IsDir() {
			return abs
		}
	}
	return "."
}

func resolveUnder(root, p string) (string, error) {
	if p == "" {
		return root, nil
	}
	if filepath.IsAbs(p) {
		return filepath.Clean(p), nil
	}
	return filepath.Abs(filepath.Join(root, p))
}

func uiURL(host string, port int) string {
	return "http://" + net.JoinHostPort(host, fmt.Sprintf("%d", port)) + "/?v=44"
}

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	proxyHost := flag.String("proxy-host", "127.0.0.1", "proxy listen address")
	proxyPort := flag.Int("proxy-port", 8080, "proxy listen port")
	apiHost := flag.String("api-host", "127.0.0.1", "API/UI listen address")
	apiPort := flag.Int("api-port", 1337, "API/UI listen port")
	callbackPort := flag.Int("callback-port", 8082, "callback receiver listen port")
	dbPath := flag.String("db-path", "aura.db", "SQLite database file")
	caDir := flag.String("ca-dir", "data", "directory for CA and leaf certificates")
	seclistsDir := flag.String("seclists-dir", "", "path to SecLists (default: <app>/third_party/SecLists)")
	openWindow := flag.Bool("window", true, "open the UI in a dedicated GTK/WebKit window")
	flag.Parse()

	addrUI := uiURL(*apiHost, *apiPort)
	if desktop.IsRunning(addrUI) {
		if !*openWindow {
			return fmt.Errorf("AURA is already running at %s", addrUI)
		}
		fmt.Fprintf(os.Stderr, "AURA already running at %s — opening window\n", addrUI)
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()
		return desktop.Open(ctx, "AURA", addrUI)
	}

	prod, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}
	defer func() { _ = prod.Sync() }()

	root := findRoot()
	if _, err := debuglog.Open(filepath.Join(root, "data", "debug")); err != nil {
		prod.Warn("debug session log disabled", zap.Error(err))
	} else {
		defer debuglog.Close()
	}
	logger := debuglog.Wrap(prod)
	if p := debuglog.Path(); p != "" {
		logger.Info("debug session log", zap.String("file", p), zap.String("latest", filepath.Join(root, "data", "debug", "latest.jsonl")))
	}
	webDir := filepath.Join(root, "web")
	caRoot, err := resolveUnder(root, *caDir)
	if err != nil {
		return fmt.Errorf("failed to resolve ca-dir: %w", err)
	}
	dbFile, err := resolveUnder(root, *dbPath)
	if err != nil {
		return fmt.Errorf("failed to resolve db path: %w", err)
	}
	if *dbPath == "aura.db" {
		if _, err := os.Stat(dbFile); err != nil {
			if legacy, lerr := resolveUnder(root, "meb.db"); lerr == nil {
				if _, err := os.Stat(legacy); err == nil {
					dbFile = legacy
				}
			}
		}
	}

	db := &storage.Store{}
	if err := db.Init(dbFile); err != nil {
		return fmt.Errorf("storage init: %w", err)
	}
	logger.Info("storage ready", zap.String("db", dbFile))

	authority := certs.NewAuthority(caRoot)
	if err := authority.Ensure(); err != nil {
		return fmt.Errorf("ca init: %w", err)
	}
	logger.Info("ca ready", zap.String("dir", caRoot))

	runtime := store.New()
	runtime.Settings.ListenHost = *proxyHost
	runtime.Settings.ListenPort = *proxyPort
	runtime.Settings.UIHost = *apiHost
	runtime.Settings.UIPort = *apiPort

	px := proxy.New(proxy.Options{
		Runtime:  runtime,
		History:  db,
		Findings: db,
		Analyzer: analyzer.DefaultEngine(),
		Certs:    authority,
		Log:      logger,
	})

	callbackSvc := callback.NewServer(*callbackPort, logger)

	batchSvc := batch.NewEngine(10, 50).WithTimeout(10 * time.Second)
	batchFactory := func(workers, rps int, onResult func(rawReq, respRaw string)) api.BatchRunner {
		return batch.NewWithLogger(workers, rps, onResult)
	}

	listsRoot := *seclistsDir
	if listsRoot == "" {
		listsRoot = filepath.Join(root, "third_party", "SecLists")
	}
	listsRoot, err = filepath.Abs(listsRoot)
	if err != nil {
		return fmt.Errorf("failed to resolve seclists-dir: %w", err)
	}
	lists := wordlist.Open(listsRoot)
	logger.Info("seclists", zap.String("dir", listsRoot), zap.Int("dictionaries", lists.Len()))

	intelEng := &intel.Engine{Store: db, Log: logger, Lists: lists}
	intelEng.Progress = func(ev intel.ProgressEvent) {
		runtime.Emit(ev)
	}

	handler := api.New(api.Options{
		Runtime:      runtime,
		Proxy:        px,
		History:      db,
		Findings:     db,
		Batch:        batchSvc,
		BatchFactory: batchFactory,
		Callback:     callbackSvc,
		Organizer:    db,
		Saver:        db,
		Intel:        intelEng,
		Lists:        lists,
		Log:          logger,
		WebDir:       webDir,
		DataDir:      caRoot,
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, ctx := errgroup.WithContext(ctx)

	if err := px.Start(ctx); err != nil {
		return err
	}

	g.Go(func() error {
		return callbackSvc.Start(ctx)
	})

	apiAddr := net.JoinHostPort(*apiHost, fmt.Sprintf("%d", *apiPort))
	httpSrv := &http.Server{
		Addr:              apiAddr,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		BaseContext:       func(net.Listener) context.Context { return ctx },
	}
	ui := uiURL(*apiHost, *apiPort)
	g.Go(func() error {
		logger.Info("api listening",
			zap.String("addr", apiAddr),
			zap.Int("proxy_port", *proxyPort),
			zap.Int("callback_port", *callbackPort),
			zap.String("db", dbFile),
			zap.String("ui", ui),
		)
		err := httpSrv.ListenAndServe()
		if err == nil || errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return fmt.Errorf("api server: %w", err)
	})
	g.Go(func() error {
		<-ctx.Done()
		sh, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = px.Stop(sh)
		if err := httpSrv.Shutdown(sh); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("api shutdown: %w", err)
		}
		return nil
	})
	if *openWindow {
		waitCtx, cancelWait := context.WithTimeout(ctx, 8*time.Second)
		err := desktop.WaitReady(waitCtx, ui)
		cancelWait()
		if err != nil {
			logger.Warn("ui window: server not ready", zap.Error(err), zap.String("url", ui))
		} else {
			logger.Info("ui window", zap.String("url", ui), zap.Bool("gtk", desktop.Native()))
			if err := desktop.Open(ctx, "AURA", ui); err != nil && !errors.Is(err, context.Canceled) {
				logger.Warn("ui window closed with error", zap.Error(err), zap.String("open_in_browser", ui))
			}
			stop()
		}
	}

	waitErr := g.Wait()

	if err := db.Close(); err != nil {
		logger.Warn("storage close", zap.Error(err))
	}
	return waitErr
}
