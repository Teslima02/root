package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/data"
	"github.com/timotew/etc/src/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TestRegister add a new User document
// Handler for HTTP Post - "/users/register"
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	// Insert User document
	cerr := repo.CreateUser(user)

	if cerr != nil {

		if strings.ContainsAny(cerr.Error(), "E11000") {
			common.DisplayAppError(
				w,
				cerr,
				"Duplicate account record!",
				http.StatusUnprocessableEntity,
			)
			return
		}

		common.DisplayAppError(
			w,
			cerr,
			cerr.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	j, err := json.Marshal(UserResource{Data: *user})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}

func Find(w http.ResponseWriter, r *http.Request) {
	var dataResource FindUsernameResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Data",
			http.StatusInternalServerError,
		)
		return
	}
	text := dataResource.Text
	inList := dataResource.InviteList
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	result, err := repo.SearchByUsername(text, inList)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"User not found",
			http.StatusNotFound,
		)
		return
	}
	if len(text) < 3 {
		common.DisplayAppError(
			w,
			errors.New("Character must be at least 3"),
			"Error invalid character length",
			http.StatusInternalServerError,
		)
		return
	}
	j, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func TempUserUpdate(w http.ResponseWriter, r *http.Request) {
	var dataResource TmpUserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Data",
			http.StatusInternalServerError,
		)
		return
	}
	context := NewContext()
	defer context.Close()
	ctx := r.Context()
	if userID := ctx.Value("user"); userID != nil {
		context.User = userID.(string)
	} else {
		common.DisplayAppError(
			w,
			errors.New("Invalid User ID"),
			"Error unrecognized context",
			http.StatusInternalServerError,
		)
		return
	}

	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	info := dataResource.Data
	ruer := &models.User{
		ID:        bson.ObjectIdHex(context.User),
		Username:  info.MatricNo,
		FirstName: info.FullName,
	}
	errv := repo.UpdateTmpUser(ruer)
	if errv != nil {
		common.DisplayAppError(
			w,
			errv,
			"User not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// Profile get user profile info
func Verify(w http.ResponseWriter, r *http.Request) {
	var token string
	context := NewContext()
	defer context.Close()
	ctx := r.Context()
	if userID := ctx.Value("user"); userID != nil {
		context.User = userID.(string)
	} else {
		common.DisplayAppError(
			w,
			errors.New("Invalid User ID"),
			"Error unrecognized context",
			http.StatusInternalServerError,
		)
		return
	}
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}

	user, err := repo.GetByID(context.User)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			http.StatusUnauthorized,
		)
		return
	}

	// Generate JWT token
	token, err = common.GenerateJWT(user.ID.Hex(), "member")
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Error while generating the access token",
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Clean-up the hashpassword to eliminate it from response JSON
	user.HashPassword = nil
	authUser := AuthUserModel{
		User:  user,
		Token: token,
	}

	j, err := json.Marshal(AuthUserResource{Data: authUser})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

// Login authenticates the HTTP request with username and apssword
// Handler for HTTP Post - "/users/login"
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			http.StatusInternalServerError,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	// Authenticate the login user
	user, err := repo.Login(loginUser)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			http.StatusUnauthorized,
		)
		return
	}
	// Generate JWT token
	token, err = common.GenerateJWT(user.ID.Hex(), "member")
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Error while generating the access token",
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Clean-up the hashpassword to eliminate it from response JSON
	user.HashPassword = nil
	authUser := AuthUserModel{
		User:  user,
		Token: token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetUserByID Handler for HTTP Get - "/users/{id}"
// Returns a single Class document by id
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	user, err := repo.GetByID(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		common.DisplayAppError(w, err, "An unexpected error has occurred", http.StatusInternalServerError)
		return

	}

	j, err := json.Marshal(user)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

// RemoveUsers remove all users
func RemoveUsers(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	err := repo.RemoveAll()
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		common.DisplayAppError(w, err, "Unable to remove all users", http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusNoContent)
}

// RemoveUser remove user permanently
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	isHex := bson.IsObjectIdHex(id)
	if !isHex {
		common.DisplayAppError(
			w,
			errors.New("Invalid Record ID"),
			"Record Not found",
			http.StatusNotFound,
		)
		return
	}
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	//Delete a test document
	err := repo.Remove(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
