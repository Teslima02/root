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

//CreateQuestion Handler for HTTP Post - "/question"
// Insert a new Note document for a TaskId
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	//get question resource
	var dataResource QuestionResource
	// Decode the incoming tEST json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Question data", 500)
		return
	}
	questionModel := dataResource.Data
	context := NewContext()
	defer context.Close()
	question := &models.Question{
		TestID: bson.ObjectIdHex(questionModel.TestID),
		Body:   questionModel.Body,
		Point:  questionModel.Point,
	}
	col := context.DbCollection("questions")
	//Insert a test document
	repo := &data.QuestionRepository{C: col}
	iErr := repo.Create(question)
	if iErr != nil {
		common.DisplayAppError(
			w,
			iErr,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	j, err := json.Marshal(question)
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

//GetQuestionsByTest Handler for HTTP Get - "/classes/tests/{id}
// Returns all tests documents under a TestID
func GetQuestionsByTest(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("questions")
	repo := &data.QuestionRepository{C: col}
	questions := repo.GetByTest(id)
	j, err := json.Marshal(QuestionsResource{Data: questions})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetQuestions Handler for HTTP Get - "/questions"
// Returns all questions documents
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("questions")
	repo := &data.QuestionRepository{C: col}
	questions, _ := repo.GetAll()
	j, err := json.Marshal(QuestionsResource{Data: questions})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetQuestionsByID Handler for HTTP Get - "/questions/{id}"
// Returns a single Question document by id
func GetQuestionByID(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("questions")
	repo := &data.QuestionRepository{C: col}
	question, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return

	}

	j, err := json.Marshal(question)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//UpdateQuestion Handler for HTTP Put - "/questions/{id}"
// Update an existing Question document
func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource QuestionResource
	// Decode the incoming Question json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Test data", 500)
		return
	}
	questionModel := dataResource.Data
	question := &models.Question{
		ID:     id,
		TestID: bson.ObjectIdHex(questionModel.TestID),
		Body:   questionModel.Body,
		Point:  questionModel.Point,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("questions")
	repo := &data.QuestionRepository{C: col}
	//Update test document
	if err := repo.Update(question); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

//DeleteQuestion Handler for HTTP Delete - "/questions/{id}"
// Delete an existing Question document
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("questions")
	repo := &data.QuestionRepository{C: col}
	//Delete a question document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
