package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
)

// ???
func NewUserDomain(
	email, password, name string, age int8,
) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
	// trocar a senha passada pela nova criptografada
}

type UserDomainInterface interface {
	// retorna caso tenha dado erro
	CreateUser() *rest_err.RestErr
	// recebe id para atualizar e retorna se deu erro ou não
	UpdateUser(string) *rest_err.RestErr
	// retorna ponteiro, pois pode estar vázio
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
