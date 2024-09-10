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
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	// recebe id para atualizar e retorna se deu erro ou não
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	// retorna ponteiro, pois pode estar vázio
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
