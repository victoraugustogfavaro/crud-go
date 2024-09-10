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
	SetID(string)
	EncryptPassword()
	GetID() string
}

// construtor
func NewUserDomain(
	email, password, name string, age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

// implementando o método get da structure |userDomain -> ud|
func (ud *userDomain) GetID() string{
	return ud.id
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

// objeto criado (privado/encapsulado)
// userDomain serve para disponibilizar e mudar o valor real
// daquele objeto, usuario
type userDomain struct {
	id       string 
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
	return ud.password
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
