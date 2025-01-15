package router

import (
	"golang_api/internal/application/contracts"
	"golang_api/internal/interface/http/handler"
	"net/http"
)

func SetupRouter(deviceService contracts.DeviceServiceContract) *http.ServeMux {
	mux := http.NewServeMux()

	deviceHandler := handler.NewDeviceHandler(deviceService)

	mux.HandleFunc("/devices", deviceHandler.CreateDevice)
	mux.HandleFunc("/devices/user", deviceHandler.GetUserDevices)

	return mux
}
