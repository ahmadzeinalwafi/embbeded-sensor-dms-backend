package handler

import (
	"context"
	aggregate "dms/internal/domain/aggregates"
	event "dms/internal/domain/events"
	tools "dms/tools"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
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

func (h *DeviceHandler) SetupDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deviceId := ps.ByName("device_id")
	var deviceConfigFields aggregate.FieldsDeviceConfig

	if !tools.DecodeJSONRequest(w, r, &deviceConfigFields) {
		return
	}

	if err := h.Validator.Struct(deviceConfigFields); err != nil {
		tools.SendErrorResponse(w, r, http.StatusBadRequest, "Validation failed", fmt.Sprintf("Validation error: %v", err))
		return
	}

	deviceConfig, err := h.DeviceService.SetupDevice(context.Background(), deviceConfigFields, deviceId)

	if err != nil {
		tools.SendErrorResponse(w, r, http.StatusInternalServerError, "Error retrieving devices", fmt.Sprintf("Error creating user: %v", err))
		return
	}

	tools.EncodeJSONResponse(w, deviceConfig, http.StatusOK)
}
func (h *DeviceHandler) GetAssosiatedUserByDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deviceId := ps.ByName("device_id")
	if deviceId == "" {
		tools.SendErrorResponse(w, r, http.StatusBadRequest, "Validation failed", "Validation error: the user_id on url query are empty")
		return
	}

	devices, err := h.DeviceService.FindAssosiatedUserByDeviceId(context.Background(), deviceId)
	if err != nil {
		tools.SendErrorResponse(w, r, http.StatusInternalServerError, "Error retrieving devices", fmt.Sprintf("Error creating user: %v", err))
		return
	}

	tools.EncodeJSONResponse(w, devices, http.StatusOK)
}

func (h *DeviceHandler) GetDeviceInfoById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deviceId := ps.ByName("device_id")
	if deviceId == "" {
		tools.SendErrorResponse(w, r, http.StatusBadRequest, "Validation failed", "Validation error: the user_id on url query are empty")
		return
	}

	devices, err := h.DeviceService.FindInfoByDeviceId(context.Background(), deviceId)
	if err != nil {
		tools.SendErrorResponse(w, r, http.StatusInternalServerError, "Error retrieving devices", fmt.Sprintf("Error creating user: %v", err))
		return
	}

	tools.EncodeJSONResponse(w, devices, http.StatusOK)
}

func (h *DeviceHandler) DeleteDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deviceId := ps.ByName("device_id")

	if deviceId == "" {
		tools.SendErrorResponse(w, r, http.StatusBadRequest, "Invalid Request", "User ID is required")
		return
	}

	if err := h.DeviceService.DeleteDevice(context.Background(), deviceId); err != nil {
		http.Error(w, fmt.Sprintf("Error deleting device: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
