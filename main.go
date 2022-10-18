package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqoi/todo-apis-go/config"
	"github.com/rifqoi/todo-apis-go/server"
	"github.com/rifqoi/todo-apis-go/server/controllers"
	gorm "github.com/rifqoi/todo-apis-go/server/repositories/gorm"
	"github.com/rifqoi/todo-apis-go/server/services"
)

func main() {

	db, err := config.ConnectPostgresGORM()
	if err != nil {
		panic(err)
	}

	todoRepo := gorm.NewTodoRepositories(db)
	todoSvc := services.NewTodoService(todoRepo)
	todoController := controllers.NewTodoController(todoSvc)

	router := gin.Default()
	app := server.NewRouter(router, todoController)

	app.Run(":4000")

}
