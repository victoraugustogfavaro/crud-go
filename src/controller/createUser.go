package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/validation"
	"github.com/victoraugustogfavaro/crud-go/src/controller/model/request"
)

// todo o contexto da requisição
func CreateUser(c *gin.Context) {

	// criando uma variável
	var userRequest request.UserRequest

	// colocar o conteúdo do body da requisição pra um objeto
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		// chamando o método
		fmt.Sprintf("Há campos incorretos\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
