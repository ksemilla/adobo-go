package auth

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ksemilla/adobo-go/pkg/users"
	"go.mongodb.org/mongo-driver/mongo"
)
type MockUsersService struct {}

func (mus *MockUsersService) FindByEmail(user *users.User, email string, filter interface{}) error {
	return errors.New("Email already in use")
}

func (mus *MockUsersService) Create(*users.SignupData) (*mongo.InsertOneResult, error) {
	return &mongo.InsertOneResult{}, nil
}

func (mus *MockUsersService) List(filter interface{}) ([]*users.User, error) {
	return nil, nil
}

func TestSignupSuccess(t *testing.T) {
	
	body := []byte(`{"email": "test@test.com", "password": "test"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(body))
	res := httptest.NewRecorder()

	handler := &AuthHandler{
		UsersService: &MockUsersService{},
	}
	handler.SignUp(res, req)
	if res.Code != http.StatusCreated {
		t.Errorf("expected success request")
	}
}

func TestSignupFail_1(t *testing.T) {
	body := []byte(`{"email": "test", "password": "test"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(body))
	res := httptest.NewRecorder()
	handler := &AuthHandler{UsersService: &MockUsersService{}}
	handler.SignUp(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected fail request for inavlid email")
	}
}

func TestSignupFail_2(t *testing.T) {
	body := []byte(`{"password": "test"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(body))
	res := httptest.NewRecorder()
	handler := &AuthHandler{UsersService: &MockUsersService{}}
	handler.SignUp(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected fail request for no email field")
	}
}

func TestSignupFail_3(t *testing.T) {
	body := []byte(`{"email": "test@test.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(body))
	res := httptest.NewRecorder()
	handler := &AuthHandler{UsersService: &MockUsersService{}}
	handler.SignUp(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected fail request for no password field")
	}
}

type MockUsersService2 struct {}

func (mus *MockUsersService2) FindByEmail(user *users.User, email string, filter interface{}) error {
	return nil
}

func TestSignupFail_4(t *testing.T) {
	body := []byte(`{"email": "test@test.com", password: "test"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(body))
	res := httptest.NewRecorder()
	handler := &AuthHandler{UsersService: &MockUsersService{}}
	handler.SignUp(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected fail request for email in use already")
	}
}