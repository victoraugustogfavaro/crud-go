package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/victoraugustogfavaro/crud-go/src/controller"
)

// iniciar as rotas -> objeto(rota) 
func InitRoutes(r *gin.RouterGroup) {
	// instanciando o caminho e passando a função que será executada nela
						// parâmetro
	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
