package models

import "users/internal/users/entity"

type Users struct {
	Name     string `bson:"name"`
	LastName string `bson:"lastName"`
	UserName string `bson:"userName"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func (model *Users) MapEntityToModel(userEntity *entity.Users) {
	model.Name = userEntity.Name
	model.LastName = userEntity.LastName
	model.UserName = userEntity.UserName
	model.Email = userEntity.Email
	model.Password = userEntity.Password
}

func (model *Users) MapEntityFromModel() *entity.Users {
	return &entity.Users{
		Name:     model.Name,
		LastName: model.LastName,
		UserName: model.UserName,
		Email:    model.Email,
		Password: model.Password,
	}
}
