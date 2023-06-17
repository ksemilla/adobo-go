package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/ksemilla/adobo-go/pkg/database"
	"github.com/ksemilla/adobo-go/pkg/users"
)

func UserRoutes(r chi.Router) {
	usersService := &users.UsersService{Col: database.GetUserCol()}
	usersResource := &users.UserHandler{
		UsersService: usersService,
	}

	r.Get("/", usersResource.List)
	r.Get("/{id}", usersResource.GetOne)
}