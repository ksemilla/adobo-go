package auth

import (
	"github.com/ksemilla/adobo-go/pkg/users"
)

type AuthHandler struct {
	usersService users.UsersService
}

func (ah *AuthHandler) SignUp() {

}