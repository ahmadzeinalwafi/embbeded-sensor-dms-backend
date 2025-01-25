package handler

import (
	"context"
	dto "dms/internal/domain/data_transfer_object"
	event "dms/internal/domain/events"
	tools "dms/tools"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	UserService event.UserService
	Validator   *validator.Validate
}

func NewUserHandler(userService event.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
		Validator:   validator.New(),
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInfo dto.EnteredUserInformation

	if !tools.DecodeJSONRequest(w, r, &userInfo) {
		return
	}

	if err := h.Validator.Struct(userInfo); err != nil {
		tools.SendErrorResponse(w, r, http.StatusBadRequest, "Validation failed", fmt.Sprintf("Validation error: %v", err))
		return
	}

	createdUser, err := h.UserService.CreateUser(context.Background(), userInfo)
	if err != nil {
		tools.SendErrorResponse(w, r, http.StatusInternalServerError, "Error creating user", fmt.Sprintf("Error creating user: %v", err))
		return
	}

	tools.EncodeJSONResponse(w, createdUser, http.StatusCreated)
}

func (h *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("user_id")

	if userID == "" {
		tools.SendErrorResponse(w, r, http.StatusBadRequest, "Invalid Request", "User ID is required")
		return
	}

	userInfo, err := h.UserService.GetUserInfo(context.Background(), userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving user info: %v", err), http.StatusInternalServerError)
		return
	}

	tools.EncodeJSONResponse(w, userInfo, http.StatusOK)
}

func (h *UserHandler) GetUserToken(w http.ResponseWriter, r *http.Request) {
	var userCredentials dto.UserCredential

	if !tools.DecodeJSONRequest(w, r, &userCredentials) {
		return
	}

	if err := h.Validator.Struct(userCredentials); err != nil {
		tools.SendErrorResponse(w, r, http.StatusBadRequest, "Validation failed", fmt.Sprintf("Validation error: %v", err))
		return
	}

	userToken, err := h.UserService.GetUserToken(context.Background(), userCredentials)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error get user token: %v", err), http.StatusInternalServerError)
		return
	}

	tools.EncodeJSONResponse(w, userToken, http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("user_id")

	if userID == "" {
		tools.SendErrorResponse(w, r, http.StatusBadRequest, "Invalid Request", "User ID is required")
		return
	}

	if err := h.UserService.DeleteUserById(context.Background(), userID); err != nil {
		http.Error(w, fmt.Sprintf("Error deleting user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
