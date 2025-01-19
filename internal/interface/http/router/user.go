package router

import (
	"dms/internal/interface/http/handler"

	"github.com/julienschmidt/httprouter"
)

// AddUserRoutes adds user-related routes to the provided router.
func AddUserRoutes(router *httprouter.Router, userHandler *handler.UserHandler) {
	router.POST("/users", WrapHandler(userHandler.CreateUser))
	router.GET("/users/info", WrapHandler(userHandler.GetUserInfo))
	router.DELETE("/users", WrapHandler(userHandler.DeleteUser))
}
