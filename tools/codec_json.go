package tools

import (
	"encoding/json"
	"net/http"
)

// DecodeJSONRequest is a utility function to decode a JSON request body into the provided struct.
// It reads the request body and attempts to unmarshal it into the given `v`.
// If an error occurs during decoding, it sends an HTTP 400 Bad Request error response with the error message.
//
// Parameters:
//   - w: The `http.ResponseWriter` to write the HTTP response.
//   - r: The `http.Request` that contains the request body to decode.
//   - v: A pointer to the struct where the decoded data will be stored.
//
// Returns:
//   - bool: Returns `true` if decoding is successful, otherwise `false` if there is an error.
func DecodeJSONRequest(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

// EncodeJSONResponse is a utility function to encode the provided data into a JSON response.
// It sets the appropriate content-type header and status code for the HTTP response,
// then attempts to encode the data as JSON. If encoding fails, an HTTP 500 Internal Server Error response is sent.
//
// Parameters:
//   - w: The `http.ResponseWriter` used to write the HTTP response.
//   - data: The data to encode into the response. This should be a Go struct, slice, or map that can be converted to JSON.
//   - statusCode: The HTTP status code to be sent with the response.
//
// Returns:
//   - None: This function does not return a value but sends the response directly to the client.
func EncodeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode) // Set the response status code

	if err := json.NewEncoder(w).Encode(map[string]interface{}{"data": data}); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}
