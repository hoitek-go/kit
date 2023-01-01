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

func ErrorBadRequest(messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Bad Request")
	}
	return Error(messages[0], http.StatusBadRequest)
}

func ErrorUnAuthorized(messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Not Logged In")
	}
	return Error(messages[0], http.StatusUnauthorized)
}

func ErrorForbidden(messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Access Denied")
	}
	return Error(messages[0], http.StatusForbidden)
}

func ErrorNotFound(messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Not Found")
	}
	return Error(messages[0], http.StatusNotFound)
}

func ErrorMethodNotAllowed(messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Method Not Allowed")
	}
	return Error(messages[0], http.StatusMethodNotAllowed)
}

func ErrorUnprocessableEntity(messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Request Temporary Blocked")
	}
	return Error(messages[0], http.StatusUnprocessableEntity)
}

func ErrorLocked(messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Request Temporary Blocked")
	}
	return Error(messages[0], http.StatusLocked)
}

func ErrorTooManyRequests(messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Too Many Requests")
	}
	return Error(messages[0], http.StatusTooManyRequests)
}

func ErrorInternalServerError(messages ...string) ErrorResponse {
	if len(messages) == 0 {
		messages = append(messages, "Something Went Wrong")
	}
	return Error(messages[0], http.StatusInternalServerError)
}
