package services

import (
	"github.com/rifqoi/todo-apis-go/server/controllers/params"
	"github.com/rifqoi/todo-apis-go/server/controllers/views"
	"github.com/rifqoi/todo-apis-go/server/repositories"
	"github.com/rifqoi/todo-apis-go/server/repositories/models"
)

type TodoService struct {
	repo repositories.TodoRepo
}

func NewTodoService(repo repositories.TodoRepo) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (svc *TodoService) CreateTodo(req *params.TodoCreate) *views.Response {
	todoModel := &models.Todo{
		Task:        req.Task,
		Description: req.Description,
		IsCompleted: req.IsCompleted,
	}
	err := svc.repo.CreateTodo(todoModel)
	if err != nil {
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessCreateResponse()
}
