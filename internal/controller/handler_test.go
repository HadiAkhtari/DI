package controller

import (
	"encoding/json"
	"myapp/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type fixedTime struct{}

func (fixedTime) Now() time.Time { return time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC) }

func TestGreet(t *testing.T) {
	h := NewHandler(service.NewGreeterService(fixedTime{}))
	mux := NewMux()
	RegisterRoutes(registerDeps{Mux: mux, Handler: h})
	for _, tc := range []struct{ query, want string }{
		{"", "Good morning myfriend"},
		{"?name=Hadi", "Good morning Hadi"},
		{"?name=%20%20", "Good morning myfriend"},
		{"?name=%20Hadi%20", "Good morning Hadi"},
	} {
		t.Run(tc.query, func(t *testing.T) {
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/greet"+tc.query, nil))
			if w.Code != http.StatusOK {
				t.Fatalf("status %d", w.Code)
			}
			if got := w.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
				t.Fatalf("content type %q", got)
			}
			var body map[string]string
			if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			if body["message"] != tc.want {
				t.Fatalf("body %v, want %q", body, tc.want)
			}
		})
	}
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/greet", nil))
	if w.Code != http.StatusMethodNotAllowed || w.Header().Get("Allow") != "GET" {
		t.Fatalf("unexpected method response: %d %v", w.Code, w.Header())
	}
}
