package routes

import (
	"github.com/NearMaick/my-first-go-app/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(request *gin.RouterGroup) {
	request.GET("/getUserById/:userId", controller.FindUserById)
	request.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	request.POST("/createUser", controller.CreateUser)
	request.PUT("/updateUser/:userId", controller.UpdateUser)
	request.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
