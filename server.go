package main

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	router http.Handler
	db *mongo.Client
}