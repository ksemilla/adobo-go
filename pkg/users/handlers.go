package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	UsersService *UsersService
}

func (uh *UserHandler) GetOne(w http.ResponseWriter, r *http.Request) {

	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	rawUser := &RawUser{}
	err = uh.UsersService.GetOne(rawUser, bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user := &User{}
	tempUser, err := json.Marshal(rawUser)
	err = json.Unmarshal(tempUser, &user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (uh *UserHandler) List(w http.ResponseWriter, r *http.Request) {

	users := &[]*RawUser{}
	err := uh.UsersService.GetList(users, map[string]string{})
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}