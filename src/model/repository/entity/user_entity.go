package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// público
type UserEntity struct {
	// ignorar caso o valor seja vázio
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Age      int8               `bson:"age,omitempty"`
}
