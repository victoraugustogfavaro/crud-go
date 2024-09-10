package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/validation"
	"github.com/victoraugustogfavaro/crud-go/src/controller/model/request"
	"github.com/victoraugustogfavaro/crud-go/src/model"
	"github.com/victoraugustogfavaro/crud-go/src/view"
	"go.uber.org/zap"
)

// todo o contexto da requisição
func (uc *userControllerInterface) CreateUser(c *gin.Context) {
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

	// pega o método da userControllerInterface
	// faz tratamento de erro e cria o usuário
	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created sucessfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}