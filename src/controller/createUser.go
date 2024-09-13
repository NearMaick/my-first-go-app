package controller

import (
	"github.com/NearMaick/my-first-go-app/src/configurations/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	err := rest_err.NewBadRequestError("Chamou essa rota errado")

	context.JSON(err.Code, err)
}
