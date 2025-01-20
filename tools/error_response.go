package tools

import (
	"encoding/json"
	"net/http"
	"time"
)

// ErrorResponse represents the standard error structure
type ErrorResponse struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

// SendErrorResponse sends a structured error response in JSON format
func SendErrorResponse(w http.ResponseWriter, r *http.Request, status int, errorType string, message string) {
	errorResponse := ErrorResponse{
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    status,
		Error:     errorType,
		Message:   message,
		Path:      r.URL.Path,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}
}
