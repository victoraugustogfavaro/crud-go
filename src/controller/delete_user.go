package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// conteúdo da requisição
func (uc *userControllerInterface) DeleteUser(c *gin.Context) {

	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	// pegando o parâmetro
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	// faz tratamento de erro e deleta o usuário no banco de dados / mongodb
	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call deleteUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"updateUser controller executed sucessfully",
		zap.String("userId", userId),
		zap.String("journey", "createUser"))

	// status, deu tudo certo
	c.Status(http.StatusOK)
}
