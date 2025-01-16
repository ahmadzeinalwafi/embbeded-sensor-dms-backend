package router

import (
	event "dms/internal/domain/events"
	"dms/internal/interface/http/handler"
	"net/http"
)

func SetupRouter(deviceService event.DeviceService) *http.ServeMux {
	mux := http.NewServeMux()

	deviceHandler := handler.NewDeviceHandler(deviceService)

	mux.HandleFunc("/devices", deviceHandler.CreateDevice)
	mux.HandleFunc("/devices/user", deviceHandler.GetUserDevices)

	return mux
}
