package router

import (
	"net/http"

	"log"
	"time"
)

func NewRouter() *http.ServeMux {
	return http.NewServeMux()
}

func WrapHandler(handlerFunc func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        handlerFunc(w, r)
        duration := time.Since(start)
        log.Printf("[INFO] %s %s\n", r.URL.Path, duration)
    })
}
