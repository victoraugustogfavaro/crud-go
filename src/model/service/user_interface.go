package service

import (
	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
	"github.com/victoraugustogfavaro/crud-go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	// retorna caso tenha dado erro
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	// recebe id para atualizar e retorna se deu erro ou não
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	// retorna ponteiro, pois pode estar vázio
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
