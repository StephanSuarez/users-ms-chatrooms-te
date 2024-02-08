package repository

import (
	"fmt"
	"users/internal/users/entity"
	"users/internal/users/repository/models"
)

func Hola() string {
	return "Hola"
}

func InsertOne(userEntity *entity.Users) error {
	userModel := models.Users{}
	userModel.MapEntityToModel(userEntity)
	fmt.Println("repository")
	fmt.Println(userModel)
	return nil
}
