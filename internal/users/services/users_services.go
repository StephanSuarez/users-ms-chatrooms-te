package services

import (
	"fmt"
	"users/internal/users/entity"
	"users/internal/users/repository"
)

func Hola_services() string {

	hola := repository.Hola()

	return hola
}

func CreateUser(userEntity *entity.Users) error {
	fmt.Println("service")
	fmt.Println(userEntity)

	repository.InsertOne(userEntity)

	return nil
}
