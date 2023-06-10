package users

import "go.mongodb.org/mongo-driver/mongo"

type SignupData struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	email string
}

type UserServiceInterface interface {
	FindByEmail(*User, string, interface{}) error
	Create(*SignupData) (*mongo.InsertOneResult, error)
}