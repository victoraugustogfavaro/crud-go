package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/validation"
	"github.com/victoraugustogfavaro/crud-go/src/controller/model/request"
	"github.com/victoraugustogfavaro/crud-go/src/model"
	"github.com/victoraugustogfavaro/crud-go/src/model/service"
	"go.uber.org/zap"
)

// todo o contexto da requisição
func CreateUser(c *gin.Context) {
	// mensagem no terminal
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	// variável do user
	var userRequest request.UserRequest

	// tratamento de erro
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	// caso passe da tratativa, chama o construtor e armazena no domain
	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	// instância o userDomainService, faz tratamento de erro e cria o usuário
	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created sucessfully",
		zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
