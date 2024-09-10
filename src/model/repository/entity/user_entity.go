package entity

import "go.mongodb.org/mongo-driver/v2/bson"

// público
type UserEntity struct {
	// ignorar caso o valor seja vázio
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
	Name     string        `bson:"name"`
	Age      int8          `bson:"age"`
}
