package main

import (
	"net/http"

	"github.com/ksemilla/adobo-go/pkg/routes"
)

func main() {
	http.ListenAndServe(":3000", routes.Router())
}