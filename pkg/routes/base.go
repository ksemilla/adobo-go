package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ksemilla/adobo-go/pkg/jwtauth"
)

func Router() http.Handler {
	tokenAuth := jwtauth.GetTokenAuth()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// PUBLIC ROUTES
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hi!, I'm currently developing this api service"))
		})
		r.Route("/api/auth", AuthRoutes)
	})

	// PRIVATE ROUTES
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier((tokenAuth)))
		r.Use(jwtauth.Authenticator)
		r.Route("/api", func(r chi.Router) {	
			r.Route("/users", UserRoutes)
		})
	})
	
	return r
}