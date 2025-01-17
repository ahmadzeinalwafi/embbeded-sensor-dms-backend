package handler

import (
	"context"
	aggregate "dms/internal/domain/aggregates"
	event "dms/internal/domain/events"
	tools "dms/tools"
	"fmt"
	"net/http"
)

type DeviceHandler struct {
	DeviceService event.DeviceService
}

func NewDeviceHandler(deviceService event.DeviceService) *DeviceHandler {
	return &DeviceHandler{
		DeviceService: deviceService,
	}
}

func (h *DeviceHandler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var userDevice aggregate.EnteredDeviceInformation

    if !tools.DecodeJSONRequest(w, r, &userDevice) {
        return
    }

	createdDevice, err := h.DeviceService.CreateDevice(context.Background(), userDevice)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating device: %v", err), http.StatusInternalServerError)
		return
	}

	tools.EncodeJSONResponse(w, createdDevice, http.StatusCreated)
}

func (h *DeviceHandler) GetUserDevices(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	devices, err := h.DeviceService.GetUserDevices(context.Background(), userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving devices: %v", err), http.StatusInternalServerError)
		return
	}

	tools.EncodeJSONResponse(w, devices, http.StatusOK)
}
