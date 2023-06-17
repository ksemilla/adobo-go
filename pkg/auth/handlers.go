package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ksemilla/adobo-go/pkg/users"
	"golang.org/x/crypto/bcrypt"
)


type AuthHandler struct {
	UsersService users.UserServiceInterface
}

func (ah *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var signupData *users.SignupData
	if err := json.NewDecoder(r.Body).Decode(&signupData);err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := validator.New().Struct(signupData); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	signupData.Password = string(hashedPassword)

	err = ah.UsersService.FindByEmail(&users.RawUser{}, signupData.Email)
	if err == nil {
		http.Error(w, "Email already in use", 400)
		return
	}
	
	res, err := ah.UsersService.Create(signupData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}
func (ah *AuthHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	// GET BODY
	var data *SignupData
	if err := json.NewDecoder(r.Body).Decode(&data);err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// GET USER BY EMAIL
	user := &users.RawUser{}
	err := ah.UsersService.FindByEmail(user, data.Email)
	if err != nil {
		fmt.Println(err)
		panic("Cannot retrieve user with given email")
	}

	// CHECK IF MATCH
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		http.Error(w, "Credentials incorrect", 400)
		return
	}

	responseToken := &ResponseToken{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseToken)
}