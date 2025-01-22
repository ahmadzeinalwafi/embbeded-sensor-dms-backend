package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewRouter initializes and returns a new HTTP router without any routes.
func NewRouter() *httprouter.Router {
	return httprouter.New()
}

func WrapHandler(handlerFunc interface{}) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		switch h := handlerFunc.(type) {
		case func(http.ResponseWriter, *http.Request, httprouter.Params):
			h(w, r, ps)
		case func(http.ResponseWriter, *http.Request):
			h(w, r)
		default:
			http.Error(w, "Invalid handler function signature", http.StatusInternalServerError)
		}
	}
}
