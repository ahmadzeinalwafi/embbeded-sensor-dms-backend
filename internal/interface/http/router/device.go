package router

import (
	"dms/internal/interface/http/handler"

	"github.com/julienschmidt/httprouter"
)

// AddDeviceRoutes adds device-related routes to the provided router.
func AddDeviceRoutes(router *httprouter.Router, deviceHandler *handler.DeviceHandler) {
	router.POST("/devices", WrapHandler(deviceHandler.CreateDevice))
	router.GET("/devices/:device_id", WrapHandler(deviceHandler.GetDeviceInfoById))
	router.DELETE("/devices/:device_id", WrapHandler(deviceHandler.DeleteDevice))
	router.GET("/devices/:device_id/user", WrapHandler(deviceHandler.GetAssosiatedUserByDevice))
	router.POST("/devices/:device_id/setup", WrapHandler(deviceHandler.SetupDevice))
}
