package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"github.com/victoraugustogfavaro/crud-go/src/controller"
	"github.com/victoraugustogfavaro/crud-go/src/controller/routes"
	"github.com/victoraugustogfavaro/crud-go/src/database/mongodb"
	"github.com/victoraugustogfavaro/crud-go/src/model/repository"
	"github.com/victoraugustogfavaro/crud-go/src/model/service"
)

func main() {
	logger.Info("About to start user application")

	// lendo as variáveis de ambiente e tratando erro
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// inicializando banco de dados e tratando o errro
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
	}

	// inicializar dependências
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	// inicializando as rotas
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	// tratamento de erros
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)

	}
}
