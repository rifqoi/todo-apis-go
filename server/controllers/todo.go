package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rifqoi/todo-apis-go/server/controllers/params"
	"github.com/rifqoi/todo-apis-go/server/controllers/views"
	"github.com/rifqoi/todo-apis-go/server/services"
)

type TodoController struct {
	svc *services.TodoService
}

func NewTodoController(svc *services.TodoService) *TodoController {
	return &TodoController{
		svc: svc,
	}
}

func (ctl *TodoController) CreateTodo(c *gin.Context) {
	var req *params.TodoCreate

	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteResponse(c, views.BadRequestResponse(err))
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		WriteResponse(c, views.BadRequestResponse(err))
		return
	}

	resp := ctl.svc.CreateTodo(req)

	WriteResponse(c, resp)
}
