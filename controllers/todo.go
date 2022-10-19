package controllers

import (
	"github.com/rifqoi/todos-api-go/controllers/params"
	responses "github.com/rifqoi/todos-api-go/responses"
	"github.com/rifqoi/todos-api-go/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllTodos godoc
// @Summary Get all todos list
// @Description Get all todos list
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} []models.Todo
// @Router /todos [get]
func GetAllTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, services.GetAllTodos())
}

// CreateTodo godoc
// @Summary Create a todo
// @Description Create a new todo with the input payload
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body params.TodoCreate true "Create todo"
// @Success 201 {object} response.Response{data=models.Todo}
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	var newTodo params.TodoCreate

	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data can not created",
		})
		return
	}

	todoResponse := services.AddNewTodo(&newTodo)

	jsonRes := responses.NewSuccessCreateResponse(todoResponse)
	c.IndentedJSON(http.StatusCreated, jsonRes)
}
