package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/validation"
	"github.com/victoraugustogfavaro/crud-go/src/controller/model/request"
	"github.com/victoraugustogfavaro/crud-go/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// conteúdo da requisição
func (uc *userControllerInterface) UpdateUser(c *gin.Context) {

	logger.Info("Init updateUser controller",
		zap.String("journey", "createUser"),
	)

	// variável do userupdate 'json'
	var userRequest request.UserUpdateRequest

	// tratamento de erro no objeto enviado
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	// pegando o parâmetro
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	// caso passe da tratativa, chama o construtor e armazena no domain
	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	// faz tratamento de erro e cria o usuário no banco de dados / mongodb
	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to call updateUser service", err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("updateUser controller executed sucessfully",
		zap.String("userId", userId),
		zap.String("journey", "createUser"))

	// status, deu tudo certo
	c.Status(http.StatusOK)
}
