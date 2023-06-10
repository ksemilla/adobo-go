package handlers

import (
	"net/http"
)

type UserHandler struct {
	service any
}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users list 2"))
}