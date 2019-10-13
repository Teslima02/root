package routers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/timotew/etc/src/controllers"
	"github.com/timotew/etc/src/data"
	"github.com/timotew/etc/src/models"
)

var token, userData, user, route = "", []byte(`{"data": {"email": "timotewpeters@gmail.com", "password": "olaoluwa"} }`), models.User{}, InitRoutes()

func TestRegister(t *testing.T) {
	context := controllers.NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	error := repo.RemoveAll()
	if error != nil {
		t.Fatal(error)
	}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	var dataResource controllers.UserResource

	req, err := http.NewRequest("POST", "/users/register", bytes.NewBuffer(userData))
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)

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
	route.ServeHTTP(_rr, _req)

	if _rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("Register didn't flag duplicate record: got %v want %v",
			_rr.Code, http.StatusUnprocessableEntity)
	}
}

func TestLogin(t *testing.T) {

	var dataResource controllers.AuthUserResource
	req, err := http.NewRequest("POST", "/users/login", bytes.NewBuffer(userData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
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
	var dataResource controllers.AuthUserResource
	req, err := http.NewRequest("GET", "/users/verify", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
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
