package controller

import (
	"context"
	"go.uber.org/fx"
	"myapp/internal/config"
	"myapp/internal/logging"
	"net/http"
	"time"
)

func RunServer(lifecycle fx.Lifecycle, mux *http.ServeMux, cfg *config.Config) {
	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logging.LogInfo("Starting server")
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logging.LogError("http server failed: " + err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			_, hasDeadline := ctx.Deadline()
			if !hasDeadline {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
				defer cancel()
			}
			return srv.Shutdown(ctx)
		},
	})
}

/*
import (
	"context"
	"myapp/internal/logging"
	"net/http"
	"time"

	"go.uber.org/fx"
	"myapp/internal/config"
)




// RunServer سرور HTTP را با Uber Fx در شروع برنامه روشن و در توقف، به‌صورت graceful خاموش می‌کند.
func RunServer(lifecycle fx.Lifecycle, mux *http.ServeMux, cfg config.Config) {
	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
		// توصیه برای پایداری و امنیت:
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logging.LogInfo("starting HTTP server")

			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {

					logging.LogError("http server failed: " + err.Error())

				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// اگر ctx مهلت ندارد، یک مهلت پیش‌فرض بدهیم:
			if _, hasDeadline := ctx.Deadline(); !hasDeadline {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
				defer cancel()
			}
			logging.LogInfo("stopping HTTP server")
			// Shutdown به‌صورت graceful: درخواست‌های درحال اجرا را تمام می‌کند.
			return srv.Shutdown(ctx)
		},
	})
}*/

/*package controller

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


*/
