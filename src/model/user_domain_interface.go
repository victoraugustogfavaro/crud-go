package model

// interface de métodos
type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	GetID() string
	SetID(string)
	EncryptPassword()
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

func NewUserUpdateDomain(
	name string, age int8,
) UserDomainInterface {
	return &userDomain{
		name:     name,
		age:      age,
	}
}