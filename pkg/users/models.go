package users

import "go.mongodb.org/mongo-driver/mongo"

type SignupData struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RawUser struct {
	Id string `json:"id" bson:"_id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id string `json:"id" bson:"_id"`
	Email string `json:"email"`
}

type MongoUser interface { User | RawUser }

type UserServiceInterface interface {
	FindByEmail(*RawUser, string) error
	Create(*SignupData) (*mongo.InsertOneResult, error)
	List(interface{}) ([]*User, error)
}