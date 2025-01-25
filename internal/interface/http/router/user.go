package router

import (
	"dms/internal/interface/http/handler"

	"net/http"
)

func AddUserRoutes(router *http.ServeMux, userHandler *handler.UserHandler) {
	router.HandleFunc("POST /users", WrapHandler(userHandler.CreateUser))
	router.HandleFunc("POST /auth/token", WrapHandler(userHandler.GetUserToken))
	router.HandleFunc("GET /users/{user_id}", WrapHandler(userHandler.GetUserInfo))
	router.HandleFunc("DELETE /users/{user_id}", WrapHandler(userHandler.DeleteUser))
}
