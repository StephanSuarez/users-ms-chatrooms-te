package http

import (
	"users/internal/users/repository"
	"users/internal/users/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type UsersDependencies struct {
	ur repository.UserRepository
	us services.UserService
	uh UserHandler
}

func NewUsersDependencies(dbInstance *mongo.Database) *UsersDependencies {
	userRepository := repository.NewUserRepository(dbInstance)
	userService := services.NewUserService(&userRepository)
	userHandler := NewUserHandler(&userService)

	return &UsersDependencies{
		ur: userRepository,
		us: userService,
		uh: userHandler,
	}
}
