package dtos

type UsersRequestDTO struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
