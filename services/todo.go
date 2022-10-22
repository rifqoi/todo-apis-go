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

func GetTodoById(id int) (*models.Todo, error) {
	for i, data := range dataTodos {
		if data.ID == id {
			return &dataTodos[i], nil
		}
	}

	return nil, errors.New("Data not found")
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

func DeleteTodoById(id int) (*models.Todo, error) {
	for i, data := range dataTodos {
		if data.ID == id {
			newData := &dataTodos

			deletedData := (*newData)[i]

			(*newData)[i] = (*newData)[len(*newData)-1]
			*newData = (*newData)[:len(*newData)-1]

			return &deletedData, nil
		}
	}

	return nil, errors.New("Data cannot be deleted")
}

func UpdateTodoById(id int, updateTodo *params.TodoUpdate) (*models.Todo, error) {
	for i, data := range dataTodos {
		if data.ID == id {
			fmt.Print(id)

			if updateTodo.Description != "" {
				*&dataTodos[i].Description = updateTodo.Description
			}

			if updateTodo.ToList != "" {
				*&dataTodos[i].ToList = updateTodo.ToList
			}
			return &dataTodos[i], nil

			// newData := &dataTodos

			// oldData := (*newData)[i]
			// updateTodo.ID = oldData.ID
			// (*newData)[i] = *updateTodo

			// return &oldData, nil
		}
	}

	return nil, errors.New("Data not found")
}
