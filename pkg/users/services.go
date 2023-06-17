package users

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersService struct {
	Col *mongo.Collection
}

func (us *UsersService) GetOne(user *RawUser, filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return us.Col.FindOne(ctx, filter).Decode(user)
}

func (us *UsersService) GetList(users *[]*RawUser ,filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := us.Col.Find(ctx, bson.D{})

	for cur.Next(ctx) {
		var user *RawUser
		
		err = cur.Decode(&user)
		if err != nil {
			break
		}
		*users = append(*users, user)
	}
	return err
}

func (us *UsersService) FindByEmail(user *RawUser, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return us.Col.FindOne(ctx, bson.M{ "email": email }).Decode(user)
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

	result := []*User{}
	for cur.Next(ctx) {
		var user *User
		
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}

func (us *UsersService) FindById(id string) {

}