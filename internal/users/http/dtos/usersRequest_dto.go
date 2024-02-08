package dtos

import (
	"users/internal/users/entity"
)

type UsersRequestDTO struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (urd *UsersRequestDTO) MapEntityToDto(userEntity *entity.Users) {
	urd.Name = userEntity.Name
	urd.LastName = userEntity.LastName
	urd.UserName = userEntity.UserName
	urd.Email = userEntity.Email
	urd.Password = userEntity.Password
}

func (urd *UsersRequestDTO) MapEntityFromDto() *entity.Users {
	return &entity.Users{
		Name:     urd.Name,
		LastName: urd.LastName,
		UserName: urd.UserName,
		Email:    urd.Email,
		Password: urd.Password,
	}
}
