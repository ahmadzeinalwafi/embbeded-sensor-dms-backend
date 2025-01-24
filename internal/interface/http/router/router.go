package router

import (
	"net/http"

	"log"
	"time"

	"github.com/julienschmidt/httprouter"
)

// NewRouter initializes and returns a new HTTP router without any routes.
func NewRouter() *httprouter.Router {
	return httprouter.New()
}

// WrapHandler wraps handlers to add logging for time spent and the request path.
func WrapHandler(handlerFunc interface{}) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Log the start time
		start := time.Now()

		// Execute the handler
		switch h := handlerFunc.(type) {
		case func(http.ResponseWriter, *http.Request, httprouter.Params):
			h(w, r, ps)
		case func(http.ResponseWriter, *http.Request):
			h(w, r)
		default:
			http.Error(w, "Invalid handler function signature", http.StatusInternalServerError)
			return
		}

		// Log the request path and time spent
		duration := time.Since(start)
		log.Printf("[INFO] %s %s\n", r.URL.Path, duration)
	}
}
