package model

import (
	"crypto/md5"
	"encoding/hex"
)

// interface de métodos
type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	EncryptPassword()
}

// construtor
func NewUserDomain(
	email, password, name string, age int8,
) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

// objeto criado (privado/encapsulado)
type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

// gets
func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.email
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}

// função de criptografar a senha
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
	// trocar a senha passada pela nova criptografada
}
