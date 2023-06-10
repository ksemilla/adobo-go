package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/ksemilla/adobo-go/handlers"
)

func UserRoutes(r chi.Router) {

	handler := &handlers.UserHandler{}

	r.Get("/", handler.GetUsers)
}