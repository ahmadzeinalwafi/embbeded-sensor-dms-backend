package handler

import (
	"context"
	aggregate "dms/internal/domain/aggregates"
	event "dms/internal/domain/events"
	tools "dms/tools"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type DeviceHandler struct {
	DeviceService event.DeviceService
	Validator     *validator.Validate
}

func NewDeviceHandler(deviceService event.DeviceService) *DeviceHandler {
	return &DeviceHandler{
		DeviceService: deviceService,
		Validator:     validator.New(),
	}
}

func (h *DeviceHandler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var userDevice aggregate.EnteredDeviceInformation

	if !tools.DecodeJSONRequest(w, r, &userDevice) {
		return
	}

	if err := h.Validator.Struct(userDevice); err != nil {
		tools.SendErrorResponse(w, r, http.StatusBadRequest, "Validation failed", fmt.Sprintf("Validation error: %v", err))
		return
	}

	createdDevice, err := h.DeviceService.CreateDevice(context.Background(), userDevice)
	if err != nil {
		tools.SendErrorResponse(w, r, http.StatusInternalServerError, "Error creating device", fmt.Sprintf("Error creating user: %v", err))
		return
	}

	tools.EncodeJSONResponse(w, createdDevice, http.StatusCreated)
}

func (h *DeviceHandler) GetUserDevices(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		tools.SendErrorResponse(w, r, http.StatusBadRequest, "Validation failed", "Validation error: the user_id on url query are empty")
		return
	}

	devices, err := h.DeviceService.GetUserDevices(context.Background(), userID)
	if err != nil {
		tools.SendErrorResponse(w, r, http.StatusInternalServerError, "Error retrieving devices", fmt.Sprintf("Error creating user: %v", err))
		return
	}

	tools.EncodeJSONResponse(w, devices, http.StatusOK)
}
