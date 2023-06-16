package auth

import (
	"encoding/json"
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

	err = ah.UsersService.FindByEmail(&users.User{}, signupData.Email, map[string]string{})
	if err == nil {
		http.Error(w, "Email already in use", 400)
		return
	}
	
	res, err := ah.UsersService.Create(signupData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}