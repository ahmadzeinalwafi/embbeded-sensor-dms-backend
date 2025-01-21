package router

import (
	"dms/internal/interface/http/handler"

	"github.com/julienschmidt/httprouter"
)

// AddDeviceRoutes adds device-related routes to the provided router.
func AddDeviceRoutes(router *httprouter.Router, deviceHandler *handler.DeviceHandler) {
	router.POST("/devices", WrapHandler(deviceHandler.CreateDevice))
	router.GET("/devices", WrapHandler(deviceHandler.GetDeviceInfoById))
	router.DELETE("/devices", WrapHandler(deviceHandler.DeleteDevice))
	router.GET("/devices/user", WrapHandler(deviceHandler.GetAssosiatedUserByDevice))
}
