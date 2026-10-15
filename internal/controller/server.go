package controller

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"go.uber.org/fx"
	"myapp/internal/config"
	"myapp/internal/logging"
)

func RunServer(lifecycle fx.Lifecycle, mux *http.ServeMux, cfg *config.Config, shutdown fx.Shutdowner) {
	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := (&net.ListenConfig{}).Listen(ctx, "tcp", cfg.Addr)
			if err != nil {
				return err
			}
			logging.LogInfo("Starting server at " + listener.Addr().String())
			go func() {
				if err := srv.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logging.LogError("http server failed: " + err.Error())
					if err := shutdown.Shutdown(fx.ExitCode(1)); err != nil {
						logging.LogError("shutdown failed: " + err.Error())
					}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				_ = srv.Close()
				return err
			}
			return nil
		},
	})
}
