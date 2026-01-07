package services

import (
	"context"
	"gin-auth/config"
	"gin-auth/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCol *mongo.Collection

func InitAuthCollection() {
	userCol = config.DB.Collection("users")
}

func CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userCol.InsertOne(ctx, user)
	return err
}

func FindByEmail(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userCol.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, err
}