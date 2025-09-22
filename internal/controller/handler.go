package controller

import (
	"encoding/json"
	"myapp/internal/service"
	"net/http"
)

type Handler struct {
	svc *service.GreeterService
}

func NewHandler(svc *service.GreeterService) *Handler {
	return &Handler{svc: svc}
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
