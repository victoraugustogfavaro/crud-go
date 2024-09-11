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

	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	// variável do user 'json'
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
	// faz tratamento de erro e cria o usuário no banco de dados / mongodb
	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call CreateUser service", err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateUser controller executed sucessfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"))

	// passar pro json que deu tudo certo, status e converte pra response
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}