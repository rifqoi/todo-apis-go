package controllers

import (
	"fmt"
	"strconv"

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

// GetTodoByID godoc
// @Summary Get todo list by id
// @Description Get todo task by id
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true  "get todo task by id"  Format(id)
// @Success 200 {array} models.Todo
// @Router /todos/{id} [get]
func GetAllTodosByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}

	todo, err := services.GetTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data Not Found",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
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

// UpdateTodo godoc
// @Summary Update a todo list by id
// @Description Update todo list by id
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "update todo task by id"
// @Success 200 {array} response.Response
// @Router /todos/update/{id} [put]
func UpdateTodo(c *gin.Context) {
	var updateTodo params.TodoUpdate

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := c.BindJSON(&updateTodo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data can not created",
		})
		return
	}
	fmt.Println(id)
	todo, errUpdate := services.UpdateTodoById(id, &updateTodo)
	if errUpdate != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data Not Found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, responses.NewSuccessResponse("Data "+todo.ToList+" updated successfully"))
}
