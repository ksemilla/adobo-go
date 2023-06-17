package database

import "go.mongodb.org/mongo-driver/mongo"

// UNUSED
type Collection[T any] interface {
	GetOne(filter interface{}) *mongo.SingleResult
	GetList(filter interface{}) []T
	Create(data interface{}) *mongo.InsertOneResult
	Update(data interface{}, filter interface{}) *mongo.UpdateResult
	Delete(filter interface{}) *mongo.DeleteResult
}