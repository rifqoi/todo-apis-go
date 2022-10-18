package views

import "net/http"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

func SuccessResponse(data interface{}) *Response {
	return &Response{
		StatusCode: http.StatusOK,
		Message:    "SUCCESS",
		Data:       data,
	}
}

func SuccessCreateResponse() *Response {
	return &Response{
		StatusCode: http.StatusCreated,
		Message:    "SUCCESS",
		Data:       "user created",
	}
}

func BadRequestResponse(err error) *Response {
	return &Response{
		StatusCode: http.StatusBadRequest,
		Message:    "BAD REQUEST",
		Error:      err.Error(),
	}
}

func InternalServerErrorResponse(err error) *Response {
	return &Response{
		StatusCode: http.StatusInternalServerError,
		Message:    "INTERNAL SERVER ERROR",
		Error:      err.Error(),
	}
}
