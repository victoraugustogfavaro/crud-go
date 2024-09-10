package service

import (
	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
	"github.com/victoraugustogfavaro/crud-go/src/model"
	"go.uber.org/zap"
)

// implementando a interface
func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	// hash da senha
	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
