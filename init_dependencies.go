package main

import (
	"github.com/victoraugustogfavaro/crud-go/src/controller"
	"github.com/victoraugustogfavaro/crud-go/src/model/repository"
	"github.com/victoraugustogfavaro/crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// inicializar dependências
func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
