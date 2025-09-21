package controller

import (
	"encoding/json"
	"go.uber.org/zap"
	"myapp/internal/service"
	"net/http"
)

type Handler struct {
	svc *service.GreeterService
	log *zap.Logger
}

func NewHandler(svc *service.GreeterService, log *zap.Logger) *Handler {
	return &Handler{svc: svc,
		log: log}
}
func (h *Handler) Greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "myfriend"
	}
	resp := map[string]string{"message": h.svc.Greet(name)}
	_ = json.NewEncoder(w).Encode(resp)
}

/*type Handler struct {
	svc *service.GreeterService
}

func NewHandler() *Handler {
	return &Handler{
		svc: service.NewGreeterService(),
	}
}

func (h *Handler) Greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "myfiend"
	}
	resp := map[string]string{"message": h.svc.Greet(name)}
	_ = json.NewEncoder(w).Encode(resp)
}*/
