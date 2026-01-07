package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json: "id"`
	NAME string `bson: "name"`
	Email  string `bson: "email"`
	Password string `bson: "-" json: "password"`
}

