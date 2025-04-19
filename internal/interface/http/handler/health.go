package handler

import (
	"net/http"

	"github.com/hellofresh/health-go/v5"
)

type HealthHandler struct {
	Health *health.Health
}

func NewHealthHandler(h *health.Health) *HealthHandler {
	return &HealthHandler{Health: h}
}

func (hh *HealthHandler) Status(w http.ResponseWriter, r *http.Request) {
	hh.Health.Handler().ServeHTTP(w, r)
}
