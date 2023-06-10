package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ksemilla/adobo-go/pkg/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil) // replace with secret key
  
	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
  }

func Router() http.Handler {
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