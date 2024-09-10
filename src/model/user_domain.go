package model

// objeto criado (privado/encapsulado)
// userDomain serve para disponibilizar e mudar o valor real do objeto, usuario
type userDomain struct {
	id       string 
	email    string 
	password string 
	name     string 
	age      int8   
}

// implementando os métodos da structure userDomain -> ud
func (ud *userDomain) GetID() string{
	return ud.id
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

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