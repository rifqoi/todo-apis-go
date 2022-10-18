package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqoi/todo-apis-go/server/controllers"
)

type Router struct {
	router *gin.Engine
	todo   *controllers.TodoController
}

func NewRouter(router *gin.Engine, todo *controllers.TodoController) *Router {
	return &Router{
		router: router,
		todo:   todo,
	}
}

func (r *Router) Run(port string) {
	r.router.POST("/todo", r.todo.CreateTodo)

	r.router.Run(port)
}
