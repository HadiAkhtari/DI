package controller

import (
	"go.uber.org/fx"
	"net/http"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

type registerDeps struct {
	fx.In

	Mux     *http.ServeMux
	Handler *Handler
}

func RegisterRoutes(d registerDeps) {
	d.Mux.HandleFunc("/greet", d.Handler.Greet)
}
