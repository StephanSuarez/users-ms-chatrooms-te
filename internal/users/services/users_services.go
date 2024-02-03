package services

import "users/internal/users/repository"

func Hola_services() string {

	hola := repository.Hola()

	return hola
}
