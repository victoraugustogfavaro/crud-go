package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/victoraugustogfavaro/crud-go/src/controller"
)

// iniciar as rotas -> objeto(rota)
func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	// instanciando o caminho e passando a função que será executada nela
	// parâmetro
	r.GET("/getUserById/:userId", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
