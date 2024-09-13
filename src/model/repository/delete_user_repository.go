package repository

import (
	"context"
	"os"

	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {

	logger.Info("Init deleteUser repository",
		zap.String("journey", "deleteUser"))

	// collection é a "tabela" em NoSQL
	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	// pegando os valores e convertendo pra entity
	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	// só vai converter o usuário se for igual o user id
	filter := bson.D{{Key: "_id", Value: userIdHex}}

	// inserir no bd
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user",
			err,
			zap.String("journey", "deleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"deleteUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	return nil
}
