package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"github.com/victoraugustogfavaro/crud-go/src/controller"
	"github.com/victoraugustogfavaro/crud-go/src/controller/routes"
	"github.com/victoraugustogfavaro/crud-go/src/model/service"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// inicializar dependências
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	// caso não tenha erros, é iniciado
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)

	}
}
