package response

import "net/http"

type HandleResponse struct {
	Fn func(w http.ResponseWriter, r *http.Request) Response
}

func (h HandleResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := h.Fn(w, r)
	switch result.(type) {
	case ErrorResponse:
		err := result.(ErrorResponse)
		JsonWithWriter(w, err, err.StatusCode)
	case SuccessResponse:
		data := result.(SuccessResponse)
		JsonWithWriter(w, data, data.StatusCode)
	}
}
