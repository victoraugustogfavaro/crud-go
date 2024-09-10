package repository

import (
	"context"
	"os"

	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
	"github.com/victoraugustogfavaro/crud-go/src/model"
	"github.com/victoraugustogfavaro/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// variável/constante global
const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser repository")

	// collection é a "tabela" em NoSQL
	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	// pegando os valores e convertendo pra entity
	value := converter.ConvertDomainToEntity(userDomain)

	// inserir
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(bson.ObjectID)
	return converter.ConvertEntityToDomain(*value), nil
}