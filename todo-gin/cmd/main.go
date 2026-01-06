package main

import (
	"log"
	"os"
	"todo-gin/config"
	"todo-gin/routes"
	"todo-gin/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.ConnectDB()
	services.InitTodoCollection()

	r := gin.Default()
	routes.TodoRoutes(r)

	port := os.Getenv("PORT")

	log.Fatal(r.Run(":" + port))
}