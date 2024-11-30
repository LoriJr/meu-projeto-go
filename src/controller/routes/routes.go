package routes

import (
	"github.com/LoriJr/meu-projeto-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup){

	r.GET("/getUserById/:userId", controller.FindUserByEmail)
	r.GET("/getUserEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

}