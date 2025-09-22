package fx

import (
	"go.uber.org/fx"
	"myapp/internal/config"
	"myapp/internal/controller"
	"myapp/internal/service"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewConfig,
			service.NewRealTime,
			service.NewGreeterService,
			controller.NewHandler,
			controller.NewMux,
		),
		fx.Invoke(
			controller.RegisterRoutes,
			controller.RunServer,
		),
	)
}
