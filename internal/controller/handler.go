package controller

import (
	"encoding/json"
	"myapp/internal/logging"
	"myapp/internal/service"
	"net/http"
	"strings"
)

type Handler struct {
	svc *service.GreeterService
}

func NewHandler(svc *service.GreeterService) *Handler {
	return &Handler{svc: svc}
}
func (h *Handler) Greet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
		return
	}
	name := strings.TrimSpace(r.URL.Query().Get("name"))
	if name == "" {
		name = "myfriend"
	}
	resp := map[string]string{"message": h.svc.Greet(name)}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logging.LogError("write greeting response: " + err.Error())
	}
}
