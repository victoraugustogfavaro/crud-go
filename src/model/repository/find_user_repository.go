package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
	"github.com/victoraugustogfavaro/crud-go/src/model"
	"github.com/victoraugustogfavaro/crud-go/src/model/repository/entity"
	"github.com/victoraugustogfavaro/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"))

	// collection é a "tabela" em NoSQL
	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	// ponteiro do UserEntity
	userEntity := &entity.UserEntity{}

	// tentar encontrar o email no banco
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	// tratamento de erros
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserbyEmail"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserbyEmail"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	// loggers com sucesso
	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserByID repository",
		zap.String("journey", "findUserByID"))

	// collection é a "tabela" em NoSQL
	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	// ponteiro do UserEntity
	userEntity := &entity.UserEntity{}

	// tentar encontrar o id no banco
	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	// tratamento de erros
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this ID: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserbyID"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by ID"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserbyEmail"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	// loggers com sucesso
	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}
