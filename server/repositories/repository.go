package repositories

import "github.com/rifqoi/todo-apis-go/server/repositories/models"

type TodoRepo interface {
	CreateTodo(todo *models.Todo) error
}
