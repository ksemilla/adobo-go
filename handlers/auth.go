package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get token"))
}

type SignupSchema struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var data SignupSchema

	if err := json.NewDecoder(r.Body).Decode(&data);err != nil {
		fmt.Println("Error")
	}

	if err := validator.New().Struct(data); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("create user"))
}