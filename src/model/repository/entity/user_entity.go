package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// público
type UserEntity struct {
	// ignorar caso o valor seja vázio
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
	Age      int8               `bson:"age"`
}
