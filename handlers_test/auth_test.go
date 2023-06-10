package handlerstest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ksemilla/adobo-go/handlers"
)

// ReverseRunes returns its argument string reversed rune-wise left to right.


func TestSignupSuccess(t *testing.T) {
	body := []byte(`{"email": "test@test.com", "password": "test"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(body))
	res := httptest.NewRecorder()
	handlers.Signup(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("expected success request")
	}
}

func TestSignupFail(t *testing.T) {
	body := []byte(`{"email": "test", "password": "test"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(body))
	res := httptest.NewRecorder()
	handlers.Signup(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected fail request")
	}

	body = []byte(`{"email": "test@test.com"}`)
	req = httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(body))
	res = httptest.NewRecorder()
	handlers.Signup(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected fail request")
	}

	body = []byte(`{"password": "test"}`)
	req = httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(body))
	res = httptest.NewRecorder()
	handlers.Signup(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected fail request")
	}
}