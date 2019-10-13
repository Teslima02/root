package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/data"
)

// import (
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"strings"

// 	"github.com/gorilla/mux"
// 	"github.com/timotew/etc/src/common"
// 	"github.com/timotew/etc/src/data"
// 	"github.com/timotew/etc/src/models"
// 	mgo "gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )

// Handler for HTTP Post - "/users/register"
func FBKitPartial(w http.ResponseWriter, r *http.Request) {
	var dataResource FBKitPartialResource
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
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	info := dataResource.Data
	user, err := repo.SearchUserByPhone(info.PhoneNumber)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"User not found",
			http.StatusNotFound,
		)
		return
	}

	if err != nil {

		common.DisplayAppError(w, err, "User not found", http.StatusNotFound)
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

func FBKitLogin(w http.ResponseWriter, r *http.Request) {
	var dataResource FBKitCompleteResource
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
	kit := &dataResource.Data
	accessToken := "AA" + "|" + common.AppConfig.FBAppID + "|" + common.AppConfig.FBAppSecret
	req, _ := http.NewRequest("GET", common.AppConfig.AKTokenExchangeBaseURL+"?code="+kit.Code+"&access_token="+accessToken+"&grant_type=authorization_code", nil)
	req.Header.Add("Accept", "application/json")
	clientOne := &http.Client{}
	resp, errThree := clientOne.Do(req)
	if errThree != nil {
		common.DisplayAppError(
			w,
			errThree,
			"Invalid User data",
			500,
		)
		return
	}
	var dataResourceOne FBKitCompleteStageOneModel
	defer resp.Body.Close()
	errOne := json.NewDecoder(resp.Body).Decode(&dataResourceOne)
	if errOne != nil {
		common.DisplayAppError(
			w,
			errOne,
			"Invalid Server Error",
			500,
		)
		return
	}

	reqTwo, _ := http.NewRequest("GET", common.AppConfig.AKEndpointBaseURL+"?access_token="+dataResourceOne.AccessToken, nil)
	req.Header.Add("Accept", "application/json")
	clientTwo := &http.Client{}
	respTwo, err := clientTwo.Do(reqTwo)

	var dataResourceTwo FBKitCompleteStageTwoResource
	defer respTwo.Body.Close()

	errTwo := json.NewDecoder(respTwo.Body).Decode(&dataResourceTwo)
	if errTwo != nil {
		common.DisplayAppError(
			w,
			errTwo,
			"Invalid Server Error",
			500,
		)
		return
	}

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	user, err := repo.SearchUserByPhone("0" + dataResourceTwo.Phone.NationalNumber)
	// Generate JWT token
	token, err := common.GenerateJWT(user.ID.Hex(), "member")
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
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(respTwo.Body)
	// newStr := buf.String()
	// fmt.Println(newStr)

}
