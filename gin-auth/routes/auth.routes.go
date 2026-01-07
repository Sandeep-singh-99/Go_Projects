package routes

import (
	"gin-auth/controllers"
	"gin-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	proctected := r.Group("/profile")
	proctected.Use(middlewares.AuthMiddleware())
	{
		proctected.GET("/", controllers.GetCurrentUser)
		proctected.POST("/logout", controllers.Logout)
	}

}