package jsonutil

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResponse writes a JSON response with the given status code.
// It sets the Content-Type header to application/json and encodes the data.
// It handles potential JSON encoding errors internally.
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	responseBody, err := json.Marshal(data) // Use json.Marshal to get bytes
	log.Printf("Data %v", responseBody)
	if err != nil {
		// Log the encoding error. This is a server-side problem.
		log.Printf("jsonutil: Failed to marshal JSON response data: %v for data: %+v", err, data)

		// Send a 500 Internal Server Error.
		// Use a specific error message for the client.
		http.Error(w, "Internal Server Error: Failed to prepare response", http.StatusInternalServerError)
		return // Crucially, return to prevent further execution
	}

	// 2. If encoding succeeded, set headers and write the pre-encoded body.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode) // Set the intended status code (e.g., 200, 201)

	_, err = w.Write(responseBody) // Write the JSON byte slice to the response body
	if err != nil {
		// This error means the data couldn't be sent over the network after headers were written.
		// It's often due to a broken pipe or client disconnecting.
		// We can only log this, as we can't change the response.
		log.Printf("jsonutil: Failed to write JSON response to client: %v", err)
	}
}

// ErrorResponse writes an error message as a JSON object with the given status code.
// The JSON object will typically have an "error" key.
func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	resp := map[string]string{"error": message}
	JSONResponse(w, resp, statusCode)
}

// MethodNotAllowed sends a 405 Method Not Allowed response.
// This is useful for `net/http` handlers that manually check r.Method.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

// BadRequest sends a 400 Bad Request response.
// Useful for general validation errors or malformed requests.
func BadRequest(w http.ResponseWriter, message string) {
	ErrorResponse(w, message, http.StatusBadRequest)
}

// NotFound sends a 404 Not Found response.
func NotFound(w http.ResponseWriter, message string) {
	ErrorResponse(w, message, http.StatusNotFound)
}

// InternalServerError sends a 500 Internal Server Error response.
// Use this for unexpected errors within your application.
func InternalServerError(w http.ResponseWriter, message string) {
	ErrorResponse(w, message, http.StatusInternalServerError)
}
