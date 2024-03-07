package http

import (
	"fmt"
	"net/http"
	"users/internal/users/http/dtos"
	"users/internal/users/services"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	us services.UserService
}

type UserHandler interface {
	CreateUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

func NewUserHandler(us *services.UserService) UserHandler {
	return &userHandler{
		us: *us,
	}
}

func (uh *userHandler) CreateUser(ctx *gin.Context) {
	userDto := dtos.UsersRequestDTO{}

	if err := ctx.ShouldBind(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Body",
		})
		return
	}

	if err := userDto.ValidateString(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uh.us.CreateUser(userDto.MapEntityFromDto()); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": userDto,
	})
}

func (uh *userHandler) GetUsers(ctx *gin.Context) {
	users, err := uh.us.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var usersResDto []dtos.UsersResDTO
	for i := 0; i < len(users); i++ {
		usersResDto = append(usersResDto, dtos.UsersResDTO(users[i]))
	}

	ctx.JSON(http.StatusOK, usersResDto)
}

func (uh *userHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uh.us.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userResDto := dtos.UsersResDTO{}
	userResDto.MapEntityToDto(user)

	ctx.JSON(http.StatusOK, userResDto)
}

func (uh *userHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	userDto := dtos.UsersRequestDTO{}
	if err := ctx.ShouldBind(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userEntity, err := uh.us.UpdateUser(id, userDto.MapEntityFromDto())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userEntity)
}

func (uh *userHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uh.us.DeleteUser(id)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user with id: %s deleted", id)})
}
