package response

import (
	"encoding/json"
	"net/http"
)

func JsonWithWriter(w http.ResponseWriter, data ErrorResponse, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}

func SuccessWithWriter(w http.ResponseWriter, data ErrorResponse) {
	JsonWithWriter(w, data, http.StatusOK)
}

func ErrorWithWriter(w http.ResponseWriter, errorResponse ErrorResponse, statusCode int, messages ...string) {
	if len(messages) > 0 {
		errorResponse["message"] = messages[0]
	}
	JsonWithWriter(w, errorResponse, statusCode)
}
