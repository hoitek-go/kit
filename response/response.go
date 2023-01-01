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

func ErrorBadRequest(w http.ResponseWriter, messages ...string) {
	if len(messages) == 0 {
		messages = append(messages, "Bad Request")
	}
	Error(w, messages[0], http.StatusBadRequest)
}

func ErrorNotFound(w http.ResponseWriter, messages ...string) {
	if len(messages) == 0 {
		messages = append(messages, "Not Found")
	}
	Error(w, messages[0], http.StatusNotFound)
}

func ErrorUnAuthorized(w http.ResponseWriter, messages ...string) {
	if len(messages) == 0 {
		messages = append(messages, "Not Logged In")
	}
	Error(w, messages[0], http.StatusUnauthorized)
}

func ErrorForbidden(w http.ResponseWriter, messages ...string) {
	if len(messages) == 0 {
		messages = append(messages, "Access Denied")
	}
	Error(w, messages[0], http.StatusForbidden)
}

func ErrorInternalServerError(w http.ResponseWriter, messages ...string) {
	if len(messages) == 0 {
		messages = append(messages, "Something Went Wrong")
	}
	Error(w, messages[0], http.StatusInternalServerError)
}

func ErrorMethodNotAllowed(w http.ResponseWriter, messages ...string) {
	if len(messages) == 0 {
		messages = append(messages, "Method Not Allowed")
	}
	Error(w, messages[0], http.StatusMethodNotAllowed)
}

func ErrorTooManyRequests(w http.ResponseWriter, messages ...string) {
	if len(messages) == 0 {
		messages = append(messages, "Too Many Requests")
	}
	Error(w, messages[0], http.StatusTooManyRequests)
}

func ErrorLocked(w http.ResponseWriter, messages ...string) {
	if len(messages) == 0 {
		messages = append(messages, "Request Temporary Blocked")
	}
	Error(w, messages[0], http.StatusLocked)
}
