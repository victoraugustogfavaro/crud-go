package service

import (
	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(
	userId string) *rest_err.RestErr {
	logger.Info("Init deleteUser model.",
		zap.String("journey", "deleteUser"))

	// tratamento de erro
	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "deleteUser"))
		return err
	}

	// logger caso dê tudo certo
	logger.Info("deleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))
	return nil
}
