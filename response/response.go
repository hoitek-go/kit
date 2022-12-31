package response

import (
	"encoding/json"
	"net/http"
)

func Send(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"statusCode": statusCode,
		"data":       data,
	})
	if err != nil {
		panic(err)
	}
}

func Success(w http.ResponseWriter, data interface{}) {
	Send(w, data, 200)
}

func Error(w http.ResponseWriter, data interface{}, statusCode int, messages ...string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errResponse := map[string]interface{}{
		"statusCode": statusCode,
		"data":       data,
	}
	if len(messages) > 0 {
		errResponse["message"] = messages[0]
	}
	err := json.NewEncoder(w).Encode(errResponse)
	if err != nil {
		panic(err)
	}
}
