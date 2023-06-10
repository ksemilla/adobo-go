package users

import "net/http"

type UserHandler struct {
	usersService any
}

func (uh *UserHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get one user"))
}