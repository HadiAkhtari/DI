package controller

import (
	"context"
	"net"
	"testing"
	"time"

	"go.uber.org/fx"
	"myapp/internal/config"
)

func TestServerRejectsOccupiedAddress(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	app := fx.New(fx.NopLogger, fx.Invoke(func(lc fx.Lifecycle, shutdown fx.Shutdowner) {
		RunServer(lc, NewMux(), &config.Config{Addr: listener.Addr().String()}, shutdown)
	}))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := app.Start(ctx); err == nil {
		_ = app.Stop(ctx)
		t.Fatal("startup succeeded with an occupied address")
	}
}
