package dtos

import (
	"users/internal/users/entity"
)

type UsersRequestDTO struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (urd *UsersRequestDTO) MapEntityToDto(userEntity *entity.Users) {
	urd.UserName = userEntity.UserName
	urd.Email = userEntity.Email
	urd.Password = userEntity.Password
}

func (urd *UsersRequestDTO) MapEntityFromDto() *entity.Users {
	return &entity.Users{
		UserName: urd.UserName,
		Email:    urd.Email,
		Password: urd.Password,
	}
}
