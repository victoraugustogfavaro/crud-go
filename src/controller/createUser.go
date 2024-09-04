package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
	"github.com/victoraugustogfavaro/crud-go/src/controller/model/request"
)

// todo o contexto da requisição
func CreateUser(c *gin.Context) {

	// criando uma variável
	var userRequest request.UserRequest
	// colocar o conteúdo do body da requisição pra um objeto
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		// chamando o método
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Há campos incorretos, error=%\n", err.Error()))
		
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}