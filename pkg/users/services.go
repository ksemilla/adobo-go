package users

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersService struct {
	Col *mongo.Collection
}

func (us *UsersService) FindByEmail(user *User, email string, filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return us.Col.FindOne(ctx, filter).Decode(user)
}

func (us *UsersService) Create(data *SignupData) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return us.Col.InsertOne(ctx, data)
}

func (us *UsersService) List(filter interface{}) ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := us.Col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var result []*User
	for cur.Next(ctx) {
		var user *User
		
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		fmt.Println("ahgehehe 1", user)
		result = append(result, user)
	}
	return result, nil
}

func (us *UsersService) FindById(id string) {

}