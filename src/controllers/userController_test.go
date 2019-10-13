package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/timotew/etc/src/data"
	"github.com/timotew/etc/src/models"
)

var token, userData, user = "", []byte(`{"data": {"email": "timotewpeters@gmail.com", "password": "olaoluwa"} }`), models.User{}

func TestRegister(t *testing.T) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	error := repo.RemoveAll()
	if error != nil {
		t.Fatal(error)
	}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	var dataResource UserResource

	req, err := http.NewRequest("POST", "/users/register", bytes.NewBuffer(userData))
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body is what we expect.
	err2 := json.NewDecoder(rr.Body).Decode(&dataResource)
	if err2 != nil {
		t.Errorf("Unexpected response body: got %v",
			rr.Body.String())
	}
	user := &dataResource.Data

	if user.Email != "timotewpeters@gmail.com" {
		t.Errorf("handler returned unexpected email: got %v want %v",
			user.Email, "timotewpeters@gmail.com")
	}

	if user.Active != false {
		t.Errorf("handler returned unexpected user active status: got %v want %v",
			user.Active, false)
	}

	if user.Deleted != false {
		t.Errorf("handler returned unexpected user deleted status: got %v want %v",
			user.Deleted, false)
	}

	// Try to check for duplicate user error
	_req, _err := http.NewRequest("POST", "/users/register", bytes.NewBuffer(userData))
	if _err != nil {
		t.Fatal(_err)
	}

	_rr := httptest.NewRecorder()
	_handler := http.HandlerFunc(Register)

	_handler.ServeHTTP(_rr, _req)

	if _rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("Register didn't flag duplicate record: got %v want %v",
			_rr.Code, http.StatusUnprocessableEntity)
	}
}

func TestLogin(t *testing.T) {

	var dataResource AuthUserResource
	req, err := http.NewRequest("POST", "/users/login", bytes.NewBuffer(userData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	err2 := json.NewDecoder(rr.Body).Decode(&dataResource)
	if err2 != nil {
		t.Errorf("Unexpected response body: got %v",
			rr.Body.String())
	}
	auth := &dataResource.Data

	if auth.User.Email != "timotewpeters@gmail.com" {
		t.Errorf("handler returned unexpected email: got %v want %v",
			auth.User.Email, "timotewpeters@gmail.com")
	}

	token = auth.Token
	user = auth.User
}

func TestVerify(t *testing.T) {
	var dataResource AuthUserResource
	req, err := http.NewRequest("GET", "/users/verify", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Verify)
	ctx := req.Context()
	ctx = context.WithValue(ctx, "user", user.ID.Hex())
	// ctx = context.WithValue(ctx, "app.user",
	// 	&YourUser{ID: "qejqjq", Email: "user@example.com"})

	req = req.WithContext(ctx)
	handler.ServeHTTP(rr, req)
	err2 := json.NewDecoder(rr.Body).Decode(&dataResource)
	if err2 != nil {
		t.Errorf("Unexpected response body: got %v",
			rr.Body.String())
	}
	auth := &dataResource.Data
	if auth.User.Email != "timotewpeters@gmail.com" {
		t.Errorf("handler returned unexpected email: got %v want %v",
			auth.User.Email, "timotewpeters@gmail.com")
	}
}
