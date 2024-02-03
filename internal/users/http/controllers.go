package http

import (
	"net/http"
	"users/internal/users/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {

	quehay := services.Hola_services()

	ctx.JSON(http.StatusOK, gin.H{
		"msg": quehay,
	})
}
