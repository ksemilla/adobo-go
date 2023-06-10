package users

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	UsersService *UsersService
}

func (uh *UserHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get one user"))
}

func (uh *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	
	users, err := uh.UsersService.List(map[string]string{})
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}