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
	"syscall"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"meb/internal/analyzer"
	"meb/internal/api"
	"meb/internal/batch"
	"meb/internal/callback"
	"meb/internal/certs"
	"meb/internal/intel"
	"meb/internal/proxy"
	"meb/internal/storage"
	"meb/internal/store"
	"meb/internal/wordlist"
)

func findRoot() string {
	var cands []string
	if wd, err := os.Getwd(); err == nil {
		cands = append(cands, wd)
	}
	if exe, err := os.Executable(); err == nil {
		cands = append(cands, filepath.Dir(exe))
	}
	for _, c := range cands {
		if st, err := os.Stat(filepath.Join(c, "web", "index.html")); err == nil && !st.IsDir() {
			return c
		}
	}
	return "."
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
	dbPath := flag.String("db-path", "meb.db", "SQLite database file")
	caDir := flag.String("ca-dir", "data", "directory for CA and leaf certificates")
	seclistsDir := flag.String("seclists-dir", "", "path to SecLists (default: <app>/third_party/SecLists)")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}
	defer func() { _ = logger.Sync() }()

	root := findRoot()
	webDir := filepath.Join(root, "web")
	caRoot, err := filepath.Abs(*caDir)
	if err != nil {
		return fmt.Errorf("failed to resolve ca-dir: %w", err)
	}
	dbFile, err := filepath.Abs(*dbPath)
	if err != nil {
		return fmt.Errorf("failed to resolve db path: %w", err)
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
		logger.Error("proxy start", zap.Error(err))
		runtime.Log("error", "could not bind proxy: "+err.Error())
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
	g.Go(func() error {
		logger.Info("api listening",
			zap.String("addr", apiAddr),
			zap.Int("proxy_port", *proxyPort),
			zap.Int("callback_port", *callbackPort),
			zap.String("db", dbFile),
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

	waitErr := g.Wait()

	if err := db.Close(); err != nil {
		logger.Warn("storage close", zap.Error(err))
	}
	return waitErr
}
