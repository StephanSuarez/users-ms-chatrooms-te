package models

import "users/internal/users/entity"

type UsersR struct {
	ID       string `bson:"_id"`
	UserName string `bson:"userName"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func (model *UsersR) MapEntityFromModel() *entity.Users {
	return &entity.Users{
		ID:       model.ID,
		UserName: model.UserName,
		Email:    model.Email,
		Password: model.Password,
	}
}
