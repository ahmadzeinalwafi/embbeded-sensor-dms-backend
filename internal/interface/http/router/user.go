package router

import (
	"dms/internal/interface/http/handler"

	"github.com/julienschmidt/httprouter"
)

// AddUserRoutes adds user-related routes to the provided router.
func AddUserRoutes(router *httprouter.Router, userHandler *handler.UserHandler) {
	router.POST("/users", WrapHandler(userHandler.CreateUser))
	router.POST("/auth/token", WrapHandler(userHandler.GetUserToken))
	router.GET("/users/:user_id/info", WrapHandler(userHandler.GetUserInfo))
	router.DELETE("/users/:user_id", WrapHandler(userHandler.DeleteUser))
}
