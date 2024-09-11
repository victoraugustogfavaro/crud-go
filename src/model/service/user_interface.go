package service

import (
	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
	"github.com/victoraugustogfavaro/crud-go/src/model"
	"github.com/victoraugustogfavaro/crud-go/src/model/repository"
)

// encapsulamento
func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	// retorna caso tenha dado erro
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	// recebe id para atualizar e retorna se deu erro ou não
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	// retorna ponteiro, pois pode estar vázio
	FindUserByIDServices(id string,) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string,) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
