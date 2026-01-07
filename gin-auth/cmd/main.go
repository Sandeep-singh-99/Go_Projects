package main

import (
	"gin-auth/config"
	"gin-auth/routes"
	"gin-auth/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.ConnectDB()
	services.InitAuthCollection()

	r := gin.Default()

	routes.AuthRoutes(r)

	log.Fatal(r.Run(":"+ os.Getenv("PORT")))
}