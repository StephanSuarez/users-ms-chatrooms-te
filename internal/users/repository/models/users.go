package models

type Users struct {
	Name     string `bson:"name"`
	LastName string `bson:"lastName"`
	UserName string `bson:"userName"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
