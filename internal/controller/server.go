package controller

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"myapp/internal/config"
	"net/http"
)

func RunServer(lifecycle fx.Lifecycle, mux *http.ServeMux, cfg config.Config, log *zap.Logger) {
	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Info("starting HTTP server", zap.String("addr", cfg.Addr))
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatal("http server startup failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}


