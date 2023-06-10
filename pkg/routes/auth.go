package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/ksemilla/adobo-go/handlers"
)

func AuthRoutes(r chi.Router) {
	r.Get("/", handlers.GetToken)
	r.Post("/signup", handlers.Signup)
}