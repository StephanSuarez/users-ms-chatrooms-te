package services

import (
	"users/internal/users/entity"
	"users/internal/users/repository"
)

func CreateUser(userEntity *entity.Users) error {
	repository.InsertOne(userEntity)

	return nil
}
