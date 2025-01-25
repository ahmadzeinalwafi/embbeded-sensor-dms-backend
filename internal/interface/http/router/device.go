package router

import (
	"dms/internal/interface/http/handler"

	"net/http"
)

func AddDeviceRoutes(router *http.ServeMux, deviceHandler *handler.DeviceHandler) {
	router.HandleFunc("POST /devices", WrapHandler(deviceHandler.CreateDevice))
	router.HandleFunc("GET /devices/{device_id}", WrapHandler(deviceHandler.GetDeviceInfoById))
	router.HandleFunc("DELETE /devices/{device_id}", WrapHandler(deviceHandler.DeleteDevice))
	router.HandleFunc("GET /devices/{device_id}/user", WrapHandler(deviceHandler.GetAssosiatedUserByDevice))
	router.HandleFunc("POST /devices/{device_id}/setup", WrapHandler(deviceHandler.SetupDevice))
	router.HandleFunc("POST /devices/{device_id}/records", WrapHandler(deviceHandler.CreateRecordsDevice))
	router.HandleFunc("GET /devices/{device_id}/records", WrapHandler(deviceHandler.ReadRecordsDevice))
}
