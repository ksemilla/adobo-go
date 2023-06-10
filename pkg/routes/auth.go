package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/ksemilla/adobo-go/pkg/auth"
	"github.com/ksemilla/adobo-go/pkg/database"
	"github.com/ksemilla/adobo-go/pkg/users"
)

func AuthRoutes(r chi.Router) {
	usersService := &users.UsersService{Col: database.GetUserCol()}
	authResource := &auth.AuthHandler{UsersService: usersService}

	r.Post("/signup", authResource.SignUp)
}