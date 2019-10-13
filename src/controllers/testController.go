package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/data"
	"github.com/timotew/etc/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetTestQuestions(w http.ResponseWriter, r *http.Request) {
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

//CreateTest Handler for HTTP Post - "/tests"
// Insert a new Note document for a TaskId
func CreateTest(w http.ResponseWriter, r *http.Request) {
	//get test resource
	var dataResource TestResource
	// Decode the incoming tEST json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Test data", 500)
		return
	}
	testModel := dataResource.Data
	context := NewContext()
	defer context.Close()
	ctx := r.Context()

	if user := ctx.Value("user"); user != nil {
		context.User = user.(string)
	}

	dur, err := time.ParseDuration(testModel.Duration)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Test Duration", 500)
		return
	}
	// use template like this
	// time.Parse("02-01-2006 15:04:05", testModel.DueDate)
	dueDate, derr := time.Parse("02-01-2006 15:04:05", testModel.DueDate)
	if derr != nil {
		common.DisplayAppError(w, derr, "Invalid Due Date data", 500)
		return
	}
	accessDate, aerr := time.Parse("02-01-2006 15:04:05", testModel.AccessDate)
	if aerr != nil {
		common.DisplayAppError(w, aerr, "Invalid Access Date data", 500)
		return
	}
	isHex := bson.IsObjectIdHex(testModel.ClassID)
	if !isHex {
		common.DisplayAppError(
			w,
			errors.New("Invalid Classroom ID"),
			"Bad classroom ID",
			http.StatusNotFound,
		)
		return
	}

	test := &models.Test{
		CreatedBy:              bson.ObjectIdHex(context.User),
		Title:                  testModel.Title,
		Description:            testModel.Description,
		ClassID:                bson.ObjectIdHex(testModel.ClassID),
		Tags:                   testModel.Tags,
		DueDate:                dueDate,
		AccessDate:             accessDate,
		Duration:               dur,
		QuestionsPerAssessment: testModel.QuestionsPerAssessment,
	}
	col := context.DbCollection("tests")
	//Insert a test document
	repo := &data.TestRepository{C: col}
	iErr := repo.Create(test)
	if iErr != nil {
		common.DisplayAppError(
			w,
			iErr,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	j, err := json.Marshal(test)
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

func GetTestsByClassKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tests")
	repo := &data.TestRepository{C: col}
	tests, _ := repo.GetByClassKey(key)
	j, err := json.Marshal(TestsResource{Data: tests})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetTestsByClass Handler for HTTP Get - "/classes/tests/{id}
// Returns all tests documents under a ClassID
func GetTestsByClass(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("tests")
	repo := &data.TestRepository{C: col}
	tests, _ := repo.GetByClass(id)
	j, err := json.Marshal(TestsResource{Data: tests})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetTestByCreator Handler for HTTP Get -"/users/tests/{id}"
// Returns all tests documents under a user
func GetTestsByCreator(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	id := Vars["id"]
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
	col := context.DbCollection("tests")
	repo := &data.TestRepository{C: col}
	tests, _ := repo.GetByCreator(id)
	j, err := json.Marshal(TestsResource{Data: tests})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetTests Handler for HTTP Get - "/tests"
// Returns all test documents
func GetTests(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
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
	defer context.Close()
	col := context.DbCollection("tests")
	repo := &data.TestRepository{C: col}
	tests, _ := repo.GetAll()
	j, err := json.Marshal(TestsResource{Data: tests})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetTestByID Handler for HTTP Get - "/tests/{id}"
// Returns a single Test document by id
func GetTestByID(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("tests")
	repo := &data.TestRepository{C: col}
	test, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return

	}

	j, err := json.Marshal(test)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//UpdateTest Handler for HTTP Put - "/tests/{id}"
// Update an existing Test document
func UpdateTest(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource TestResource
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Test data", 500)
		return
	}
	testModel := dataResource.Data

	dur, err := time.ParseDuration(testModel.Duration)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Test Duration", 500)
		return
	}
	dueDate, derr := time.Parse("02-01-2006 15:04:05", testModel.DueDate)
	if derr != nil {
		common.DisplayAppError(w, derr, "Invalid Due Date data", 500)
		return
	}
	accessDate, aerr := time.Parse("02-01-2006 15:04:05", testModel.AccessDate)
	if aerr != nil {
		common.DisplayAppError(w, aerr, "Invalid Access Date data", 500)
		return
	}
	test := &models.Test{
		ID:          id,
		Title:       testModel.Title,
		Description: testModel.Description,
		ClassID:     bson.ObjectIdHex(testModel.ClassID),
		Tags:        testModel.Tags,
		DueDate:     dueDate,
		AccessDate:  accessDate,
		Duration:    dur,
	}

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tests")
	repo := &data.TestRepository{C: col}
	//Update test document
	if err := repo.Update(test); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

//DeleteTest Handler for HTTP Delete - "/tests/{id}"
// Delete an existing Test document
func DeleteTest(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("tests")
	repo := &data.TestRepository{C: col}
	//Delete a test document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
