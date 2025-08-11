package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Payload any    `json:"payload,omitempty"`
}

func JsonResponse(w http.ResponseWriter, statusCode int, success bool, message string, data any) {
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusText(statusCode),
		Success: success,
		Message: message,
		Payload: data,
	})
}
