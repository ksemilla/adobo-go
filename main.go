package main

import (
	"net/http"

	"github.com/ksemilla/adobo-go/pkg/routes"
)

func main() {
	// db := CreateDBClient()

	s := &Server{
		// db: db,
		router: routes.Router(),
	}



	http.ListenAndServe(":3000", s.router)
}