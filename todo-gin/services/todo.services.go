package services

import (
	"context"
	"time"
	"todo-gin/config"
	"todo-gin/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoCollection *mongo.Collection

func InitTodoCollection() {
	todoCollection = config.DB.Collection("todos")
}

func CreateTodo(todo models.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := todoCollection.InsertOne(ctx, todo)
	return err
}

func GetTodos() ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := todoCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var todos []models.Todo 
	err = cursor.All(ctx, &todos)
	return todos, err 
}

func DeleteTodo(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err := todoCollection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}