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

//CreateChoice Handler for HTTP Post - "/choice"
// Insert a new Choice document
func CreateChoice(w http.ResponseWriter, r *http.Request) {
	//get choice resource
	var dataResource ChoiceResource
	// Decode the incoming Choice json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Choice data", 500)
		return
	}
	choiceModel := dataResource.Data
	context := NewContext()
	defer context.Close()
	choice := &models.Choice{
		QuestionID: bson.ObjectIdHex(choiceModel.QuestionID),
		Correct:    choiceModel.Correct,
		OptionID:   bson.ObjectIdHex(choiceModel.OptionID),
	}
	col := context.DbCollection("choices")
	//Insert a choice document
	repo := &data.ChoiceRepository{C: col}
	iErr := repo.Create(choice)
	if iErr != nil {
		common.DisplayAppError(
			w,
			iErr,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	j, err := json.Marshal(choice)
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

//GetChoicesByQuestion Handler for HTTP Get - "/choices/questions/{id}
// Returns all tests documents under a QuestionID
func GetChoicesByQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("choices")
	repo := &data.ChoiceRepository{C: col}
	choices := repo.GetByQuestion(id)
	j, err := json.Marshal(ChoicesResource{Data: choices})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetChoices Handler for HTTP Get - "/choices"
// Returns all choices documents
func GetChoices(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("choices")
	repo := &data.ChoiceRepository{C: col}
	choices, _ := repo.GetAll()
	j, err := json.Marshal(ChoicesResource{Data: choices})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetChoiceByID Handler for HTTP Get - "/choices/{id}"
// Returns a single Choice document by id
func GetChoiceByID(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("choices")
	repo := &data.ChoiceRepository{C: col}
	choice, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return

	}

	j, err := json.Marshal(choice)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//UpdateChoice Handler for HTTP Put - "/choices/{id}"
// Update an existing Choice document
func UpdateChoice(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	isHex := bson.IsObjectIdHex(vars["id"])
	if !isHex {
		common.DisplayAppError(
			w,
			errors.New("Invalid Record ID"),
			"Record Not found",
			http.StatusNotFound,
		)
		return
	}
	id := bson.ObjectIdHex(vars["id"])
	var dataResource ChoiceResource
	// Decode the incoming Choice json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Choice data", 500)
		return
	}
	choiceModel := dataResource.Data
	choice := &models.Choice{
		ID:         id,
		QuestionID: bson.ObjectIdHex(choiceModel.QuestionID),
		Correct:    choiceModel.Correct,
		OptionID:   bson.ObjectIdHex(choiceModel.OptionID),
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("choices")
	repo := &data.ChoiceRepository{C: col}
	//Update test document
	if err := repo.Update(choice); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

//DeleteChoice Handler for HTTP Delete - "/choices/{id}"
// Delete an existing Choice document
func DeleteChoice(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("choices")
	repo := &data.ChoiceRepository{C: col}
	//Delete a choice document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
