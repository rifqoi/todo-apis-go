package response

import (
	"net/http"

	"github.com/rifqoi/todos-api-go/services/models"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewSuccessCreateResponse(data *models.Todo) *Response {
	return &Response{
		StatusCode: http.StatusCreated,
		Message:    "data created successfully",
		Data:       *data,
	}
}

func NewSuccessResponse(msg string) *Response {
	return &Response{
		StatusCode: http.StatusOK,
		Message:    msg,
	}
}
