package routes

import (
	"todo-gin/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine) {
	todo := r.Group("/api/todos")
	{
		todo.POST("/", controllers.CreateTodo)
		todo.GET("/", controllers.GetTodos)
		todo.DELETE("/:id", controllers.DeleteTodo)
	}
}