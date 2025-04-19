package router

import (
	"dms/internal/interface/http/handler"

	"net/http"
)

func AddHealthCheckRoutes(router *http.ServeMux, healthCheckHandler *handler.HealthHandler) {
	router.HandleFunc("GET /health", WrapHandler(healthCheckHandler.Status))
}
