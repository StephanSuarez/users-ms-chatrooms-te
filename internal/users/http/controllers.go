package http

import (
	"net/http"
	"users/internal/users/http/dtos"
	"users/internal/users/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {

	quehay := services.Hola_services()

	ctx.JSON(http.StatusOK, gin.H{
		"msg": quehay,
	})
}

func AddUser(ctx *gin.Context) {

	userDto := dtos.UsersRequestDTO{}

	if err := ctx.ShouldBind(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad body",
		})
		return
	}

	if err := services.CreateUser(userDto.MapEntityFromDto()); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": userDto,
	})
}
