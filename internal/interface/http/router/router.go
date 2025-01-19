package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewRouter initializes and returns a new HTTP router without any routes.
func NewRouter() *httprouter.Router {
	return httprouter.New()
}

func WrapHandler(handlerFunc http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handlerFunc(w, r)
	}
}
