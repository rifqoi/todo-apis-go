package routes

import (
	"github.com/rifqoi/todos-api-go/controllers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func TodoRouters(router *gin.Engine) {
	router.GET("/todos", controllers.GetAllTodos)
	router.GET("/todos/:id", controllers.GetAllTodosByID)
	router.POST("/todos", controllers.CreateTodo)
	router.PUT("/todos/update/:id", controllers.UpdateTodo)
	router.DELETE("/todos/delete/:id", controllers.DeleteTodo)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
