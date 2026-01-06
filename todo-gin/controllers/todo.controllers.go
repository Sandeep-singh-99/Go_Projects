package controllers

import (
	"net/http"
	"todo-gin/models"
	"todo-gin/services"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.Completed = false

	err := services.CreateTodo(todo)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Todo created successfully"})
}

func GetTodos(c *gin.Context) {
	todos, err := services.GetTodos()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, todos)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteTodo(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Todo deleted successfully"})
}