package services

import (
	"errors"
	"users/internal/users/entity"
	"users/internal/users/repository"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	ur repository.UserRepository
}

type UserService interface {
	CreateUser(userEntity *entity.Users) error
	GetUsers() ([]entity.Users, error)
	GetUserByID(id string) (*entity.Users, error)
	UpdateUser(id string, userEntity *entity.Users) (*entity.Users, error)
	DeleteUser(id string) error
}

func NewUserService(ur *repository.UserRepository) UserService {
	return &userService{
		ur: *ur,
	}
}

func (us *userService) CreateUser(userEntity *entity.Users) error {
	hashPass, err := hashPassword(userEntity.Password)
	if err != nil {
		return err
	}
	userEntity.Password = hashPass
	if err = us.ur.InsertOne(userEntity); err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (us *userService) GetUsers() ([]entity.Users, error) {
	users, err := us.ur.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userService) GetUserByID(id string) (*entity.Users, error) {
	user, err := us.ur.FindOne(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) UpdateUser(id string, userEntity *entity.Users) (*entity.Users, error) {
	user, err := us.ur.UpdateOne(id, userEntity)
	if err != nil {
		return nil, err
	}
	return user, nil
}

var ErrUserNotFound = errors.New("user ID was not found")

func (us *userService) DeleteUser(id string) error {
	userDeleted, err := us.ur.DeleteOne(id)
	if !userDeleted {
		return ErrUserNotFound
	}

	if err != nil {
		return err
	}

	return nil
}
