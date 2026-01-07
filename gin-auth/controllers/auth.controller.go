package controllers

import (
	"gin-auth/models"
	"gin-auth/services"
	"gin-auth/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var body models.User
	c.BindJSON(&body)

	hash, _ := utils.HashPassword(body.Password)
	body.Password = hash

	err := services.CreateUser(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "User Exists"})
		return
	}

	c.JSON(200, gin.H{"message": "User Created"})
}

func Login(c *gin.Context) {
	var body models.User
	c.BindJSON(&body)

	user, err := services.FindByEmail(body.Email)
	if err != nil || !utils.CheckPassword(user.Password, body.Password) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID.Hex())

	c.SetCookie(
		"token", token, 3600*24, "/", "", false, true,
	)

	c.JSON(200, gin.H{"message": "Logged In"})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{"message": "Logged Out"})
}


func GetCurrentUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "User not found in context"})
		return
	}

	c.JSON(200, gin.H{
		"userId": user,
	})
}