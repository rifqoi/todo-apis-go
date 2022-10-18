package services

import (
	"github.com/rifqoi/todo-apis-go/server/repositories"
	"github.com/rifqoi/todo-apis-go/server/repositories/models"
	"gorm.io/gorm"
)

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepositories(db *gorm.DB) repositories.TodoRepo {
	return &todoRepo{
		db: db,
	}
}

func (repo *todoRepo) CreateTodo(todo *models.Todo) error {
	err := repo.db.Create(&todo).Error
	return err
}
