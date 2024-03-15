package dtos

import (
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/entity"
)

type UsersResDTO struct {
	ID       string
	UserName string
	Email    string
	Password string
}

func (urd *UsersResDTO) MapEntityToDto(userEntity *entity.Users) {
	urd.ID = userEntity.ID
	urd.UserName = userEntity.UserName
	urd.Email = userEntity.Email
	urd.Password = userEntity.Password
}

func (urd *UsersResDTO) MapEntityFromDto() *entity.Users {
	return &entity.Users{
		ID:       urd.ID,
		UserName: urd.UserName,
		Email:    urd.Email,
		Password: urd.Password,
	}
}
