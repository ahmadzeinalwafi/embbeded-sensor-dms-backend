package router

import (
	event "dms/internal/domain/events"
	"dms/internal/interface/http/handler"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SetupRouter initializes the HTTP router with routes for device handling.
func SetupRouter(deviceService event.DeviceService) *httprouter.Router {
	router := httprouter.New()

	// Initialize the device handler with the provided service.
	deviceHandler := handler.NewDeviceHandler(deviceService)

	// Define routes and map them to the appropriate handler methods.
	router.POST("/devices", wrapHandler(deviceHandler.CreateDevice))
	router.GET("/devices/user", wrapHandler(deviceHandler.GetUserDevices))

	return router
}

// wrapHandler converts http.HandlerFunc to httprouter.Handle.
func wrapHandler(handlerFunc http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handlerFunc(w, r)
	}
}
