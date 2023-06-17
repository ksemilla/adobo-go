package users_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ksemilla/adobo-go/pkg/users"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockUsersService struct {}

func (mus *MockUsersService) FindByEmail(user *users.RawUser, email string) error {
	return errors.New("Email already in use")
}

func (mus *MockUsersService) Create(*users.SignupData) (*mongo.InsertOneResult, error) {
	return &mongo.InsertOneResult{}, nil
}

func (mus *MockUsersService) List(filter interface{}) ([]*users.User, error) {
	return nil, nil
}

func (mus *MockUsersService) GetList(users *[]*users.RawUser ,filter interface{}) error {
	return nil
}

func (mus *MockUsersService) GetOne(users *users.RawUser, filter interface{}) error {
	return nil
}

func TestGetOneFail(t *testing.T) {
	
	req := httptest.NewRequest(http.MethodGet, "/api/users/", nil)
	res := httptest.NewRecorder()

	handler := &users.UserHandler{
		UsersService: &MockUsersService{},
	}
	handler.GetOne(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected success request")
	}
}