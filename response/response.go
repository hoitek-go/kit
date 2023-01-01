package response

import (
	"net/http"
)

type ErrorResponse = map[string]interface{}

func BuildErrorResponse(data interface{}, statusCode int) ErrorResponse {
	return ErrorResponse{
		"statusCode": statusCode,
		"data":       data,
	}
}

func Success(data interface{}) ErrorResponse {
	return BuildErrorResponse(data, 200)
}

func Error(data interface{}, statusCode int, messages ...string) ErrorResponse {
	errResponse := BuildErrorResponse(data, statusCode)
	if len(messages) > 0 {
		errResponse["message"] = messages[0]
	}
	return errResponse
}

func ErrorBadRequest(data interface{}, messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Bad Request")
	}
	return Error(data, http.StatusBadRequest, messages[0])
}

func ErrorUnAuthorized(data interface{}, messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Not Logged In")
	}
	return Error(data, http.StatusUnauthorized, messages[0])
}

func ErrorForbidden(data interface{}, messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Access Denied")
	}
	return Error(data, http.StatusForbidden, messages[0])
}

func ErrorNotFound(data interface{}, messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Not Found")
	}
	return Error(data, http.StatusNotFound, messages[0])
}

func ErrorMethodNotAllowed(data interface{}, messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Method Not Allowed")
	}
	return Error(data, http.StatusMethodNotAllowed, messages[0])
}

func ErrorUnprocessableEntity(data interface{}, messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Request Temporary Blocked")
	}
	return Error(data, http.StatusUnprocessableEntity, messages[0])
}

func ErrorLocked(data interface{}, messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Request Temporary Blocked")
	}
	return Error(data, http.StatusLocked, messages[0])
}

func ErrorTooManyRequests(data interface{}, messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Too Many Requests")
	}
	return Error(data, http.StatusTooManyRequests, messages[0])
}

func ErrorInternalServerError(data interface{}, messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Something Went Wrong")
	}
	return Error(data, http.StatusInternalServerError, messages[0])
}
