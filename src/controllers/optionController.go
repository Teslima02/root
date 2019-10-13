package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/data"
	"github.com/timotew/etc/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//CreateOption Handler for HTTP Post - "/option"
// Insert a new Note document for a TaskId
func CreateOption(w http.ResponseWriter, r *http.Request) {
	//get option resource
	var dataResource OptionResource
	// Decode the incoming tEST json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Option data", 500)
		return
	}
	optionModel := dataResource.Data
	context := NewContext()
	defer context.Close()
	isHex := bson.IsObjectIdHex(optionModel.QuestionID)

	if !isHex {
		common.DisplayAppError(
			w,
			errors.New("Invalid Question ID"),
			"Bad question ID",
			http.StatusNotFound,
		)
		return
	}
	option := &models.Option{
		Correct:    optionModel.Correct,
		QuestionID: bson.ObjectIdHex(optionModel.QuestionID),
		Body:       optionModel.Body,
	}
	col := context.DbCollection("options")
	//Insert a option document
	repo := &data.OptionRepository{C: col}
	iErr := repo.Create(option)
	if iErr != nil {
		common.DisplayAppError(
			w,
			iErr,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	j, err := json.Marshal(option)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//GetOptionsByQuestion Handler for HTTP Get - "/options/questions/{id}
// Returns all tests documents under a QuestionID
func GetOptionsByQuestion(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("options")
	repo := &data.OptionRepository{C: col}
	options, _ := repo.GetByQuestion(id)
	j, err := json.Marshal(OptionsResource{Data: options})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetOptions Handler for HTTP Get - "/options"
// Returns all options documents
func GetOptions(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("options")
	repo := &data.OptionRepository{C: col}
	options, _ := repo.GetAll()
	j, err := json.Marshal(OptionsResource{Data: options})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetOptionByID Handler for HTTP Get - "/options/{id}"
// Returns a single Option document by id
func GetOptionByID(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("options")
	repo := &data.OptionRepository{C: col}
	option, err := repo.GetByID(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return

	}

	j, err := json.Marshal(option)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//UpdateOption Handler for HTTP Put - "/options/{id}"
// Update an existing Option document
func UpdateOption(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource OptionResource
	// Decode the incoming Option json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Option data", 500)
		return
	}
	optionModel := dataResource.Data
	option := &models.Option{
		ID:         id,
		QuestionID: bson.ObjectIdHex(optionModel.QuestionID),
		Body:       optionModel.Body,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("options")
	repo := &data.OptionRepository{C: col}
	//Update test document
	if err := repo.Update(option); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

//DeleteOption Handler for HTTP Delete - "/options/{id}"
// Delete an existing Option document
func DeleteOption(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("options")
	repo := &data.OptionRepository{C: col}
	//Delete a option document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
