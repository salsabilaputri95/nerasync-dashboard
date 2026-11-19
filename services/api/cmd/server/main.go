package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"nerasync-dashboard/services/api/internal/config"
	apihttp "nerasync-dashboard/services/api/internal/http"
	"nerasync-dashboard/services/api/internal/store"
)

func main() {
	cfg := config.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: parseLogLevel(cfg.LogLevel),
	}))

	rootCtx := context.Background()
	publicPool, err := store.NewPublicPool(rootCtx, cfg.DatabaseURLPublic)
	if err != nil {
		logger.Error("failed to initialize public database pool", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer publicPool.Close()

	router := apihttp.NewRouter(cfg.AppEnv, cfg.CORSAllowedOrigins, publicPool)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("sinergi-api listening", slog.String("port", cfg.Port), slog.String("env", cfg.AppEnv))
		if serveErr := srv.ListenAndServe(); serveErr != nil && !errors.Is(serveErr, http.ErrServerClosed) {
			errCh <- serveErr
		}
		close(errCh)
	}()

	sigCtx, stop := signal.NotifyContext(rootCtx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case serveErr := <-errCh:
		if serveErr != nil {
			logger.Error("server exited unexpectedly", slog.String("error", serveErr.Error()))
			os.Exit(1)
		}
	case <-sigCtx.Done():
		logger.Info("shutdown signal received, draining connections")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
		defer cancel()
		if shutdownErr := srv.Shutdown(shutdownCtx); shutdownErr != nil {
			logger.Error("graceful shutdown failed", slog.String("error", shutdownErr.Error()))
			os.Exit(1)
		}
		logger.Info("server shutdown complete")
	}
}

func parseLogLevel(raw string) slog.Level {
	switch raw {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
