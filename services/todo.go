package services

import (
	"errors"
	"fmt"
	"github.com/rifqoi/todos-api-go/controllers/params"
	"github.com/rifqoi/todos-api-go/services/models"
)

var dataTodos = []models.Todo{
	{ID: 1, ToList: "Belanja", Description: "Belanja Bulanan"},
	{ID: 2, ToList: "Seminar", Description: "Seminar nasional di kampus"},
	{ID: 3, ToList: "dokter hewan", Description: "Memeriksa kucing"},
}

func GetAllTodos() []models.Todo {
	return dataTodos
}

func AddNewTodo(newTodo *params.TodoCreate) *models.Todo {
	data := &dataTodos
	newId := (*data)[len(*data)-1].ID + 1
	// newTodo.ID = newId
	todo := models.Todo{
		ID:          newId,
		ToList:      newTodo.ToList,
		Description: newTodo.Description,
	}
	dataTodos = append(dataTodos, todo)
	return &todo
}
