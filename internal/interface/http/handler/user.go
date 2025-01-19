package handler

import (
	"context"
	aggregate "dms/internal/domain/aggregates"
	event "dms/internal/domain/events"
	tools "dms/tools"
	"fmt"
	"net/http"
)

type UserHandler struct {
	UserService event.UserService
}

func NewUserHandler(userService event.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInfo aggregate.EnteredUserInformation

	if !tools.DecodeJSONRequest(w, r, &userInfo) {
		return
	}

	createdUser, err := h.UserService.CreateUser(context.Background(), userInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
		return
	}

	tools.EncodeJSONResponse(w, createdUser, http.StatusCreated)
}

func (h *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	userInfo, err := h.UserService.GetUserInfo(context.Background(), userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving user info: %v", err), http.StatusInternalServerError)
		return
	}

	tools.EncodeJSONResponse(w, userInfo, http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	if err := h.UserService.DeleteUserById(context.Background(), userID); err != nil {
		http.Error(w, fmt.Sprintf("Error deleting user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
