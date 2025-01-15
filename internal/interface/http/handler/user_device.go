package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_api/internal/application/contracts"
	"net/http"
)

type DeviceHandler struct {
	DeviceService contracts.DeviceServiceContract
}

func NewDeviceHandler(deviceService contracts.DeviceServiceContract) *DeviceHandler {
	return &DeviceHandler{
		DeviceService: deviceService,
	}
}

func (h *DeviceHandler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var userDevice contracts.EnteredDeviceInformation

	// Decode the incoming request body into the device data structure
	if err := json.NewDecoder(r.Body).Decode(&userDevice); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// Call the use case to create the device
	createdDevice, err := h.DeviceService.CreateDevice(context.Background(), userDevice)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating device: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdDevice)
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(devices)
}
