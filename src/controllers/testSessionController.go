package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/data"
	"github.com/timotew/etc/src/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CloseTestSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	isHex := bson.IsObjectIdHex(id)
	fmt.Println(id)
	if !isHex {

		common.DisplayAppError(
			w,
			errors.New("Invalid Record ID"),
			"Record Not found",
			http.StatusNotFound,
		)
		return
	}
	ctx := r.Context()

	if user := ctx.Value("user"); user != nil {
		context.User = user.(string)
	}
	col := context.DbCollection("testSessions")
	repo := &data.TestSessionRepository{C: col}
	err := repo.CloseSession(id, context.User)
	if err != nil {

		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}

func GetTestSessionByTestID(w http.ResponseWriter, r *http.Request) {
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
	ctx := r.Context()

	if user := ctx.Value("user"); user != nil {
		context.User = user.(string)
	}
	col := context.DbCollection("testSessions")
	repo := &data.TestSessionRepository{C: col}
	session, err := repo.GetByTestId(id, context.User)
	if err != nil {
		if err != mgo.ErrNotFound {
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}

		// get test
		testCol := context.DbCollection("tests")
		testRepo := &data.TestRepository{C: testCol}
		test, tErr := testRepo.GetById(id)
		if tErr != nil {
			common.DisplayAppError(
				w,
				tErr,
				"Invalid Test",
				500,
			)
			return
		}

		// get question
		questionCol := context.DbCollection("questions")
		questionRepo := &data.QuestionRepository{C: questionCol}
		questionIds, rErr := questionRepo.GetRandomQuestionsIDsByTestID(id, test.QuestionsPerAssessment)
		if rErr != nil {
			common.DisplayAppError(
				w,
				rErr,
				"An unexpected error has occurred",
				500,
			)
			return
		}

		bsonIds := []bson.ObjectId{}
		for _, bid := range questionIds {
			bsonIds = append(bsonIds, bid.ID)
		}
		// create session and send
		session := &models.TestSession{
			TestID:    bson.ObjectIdHex(id),
			UserID:    bson.ObjectIdHex(context.User),
			Questions: bsonIds,
			Time:      test.Duration,
		}

		col := context.DbCollection("testSessions")
		//Insert a test document
		repo := &data.TestSessionRepository{C: col}
		iErr := repo.Create(session)
		if iErr != nil {
			common.DisplayAppError(
				w,
				iErr,
				"An unexpected error has occurred",
				500,
			)
			return
		}

		j, err := json.Marshal(session)
		if err != nil {
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		return
	}

	j, err := json.Marshal(session)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

func GetTestSessionByID(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("testSessions")
	repo := &data.TestSessionRepository{C: col}
	test, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
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

func CreateTestSession(w http.ResponseWriter, r *http.Request) {
	//get question resource
	var dataResource TestSessionResource
	// Decode the incoming tEST json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Session data", 500)
		return
	}
	testSessionModel := dataResource.Data
	context := NewContext()
	defer context.Close()
	testIsHex := bson.IsObjectIdHex(testSessionModel.TestID)
	if !testIsHex {
		common.DisplayAppError(
			w,
			errors.New("Invalid Test ID"),
			"Test Not found",
			http.StatusNotFound,
		)
		return
	}

	ctx := r.Context()
	if user := ctx.Value("user"); user != nil {
		context.User = user.(string)
	}

	// get test
	testCol := context.DbCollection("tests")
	testRepo := &data.TestRepository{C: testCol}
	test, tErr := testRepo.GetById(testSessionModel.TestID)
	if tErr != nil {
		common.DisplayAppError(
			w,
			tErr,
			"Invalid Test",
			500,
		)
		return
	}
	// get question
	questionCol := context.DbCollection("questions")
	questionRepo := &data.QuestionRepository{C: questionCol}
	questionIds, rErr := questionRepo.GetRandomQuestionsIDsByTestID(testSessionModel.TestID, test.QuestionsPerAssessment)
	if rErr != nil {
		common.DisplayAppError(
			w,
			rErr,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	bsonIds := []bson.ObjectId{}
	for _, bid := range questionIds {
		bsonIds = append(bsonIds, bid.ID)
	}

	session := &models.TestSession{
		TestID:    bson.ObjectIdHex(testSessionModel.TestID),
		UserID:    bson.ObjectIdHex(context.User),
		Questions: bsonIds,
		Time:      test.Duration,
	}
	col := context.DbCollection("testSessions")
	//Insert a test document
	repo := &data.TestSessionRepository{C: col}
	iErr := repo.Create(session)
	if iErr != nil {
		common.DisplayAppError(
			w,
			iErr,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	j, err := json.Marshal(session)
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

func UpdateTestSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource TestSessionResource
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Session data", 500)
		return
	}
	testSessionModel := dataResource.Data

	time, err := time.ParseDuration(testSessionModel.Time)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Time", 500)
		return
	}
	session := &models.TestSession{
		ID:     id,
		Time:   time,
		Active: testSessionModel.Active,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("testSessions")
	//Insert a test document
	repo := &data.TestSessionRepository{C: col}
	iErr := repo.Update(session)
	if iErr != nil {
		common.DisplayAppError(
			w,
			iErr,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetTestSessionQuestions
func GetTestSessionQuestions(w http.ResponseWriter, r *http.Request) {
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
	col := context.DbCollection("testSessions")
	repo := &data.TestSessionRepository{C: col}
	testSession, err := repo.GetById(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Test",
			500,
		)
		return
	}
	questionCol := context.DbCollection("questions")
	questionRepo := &data.QuestionRepository{C: questionCol}
	questions, rErr := questionRepo.GetManyByObjectId(testSession.Questions)
	if rErr != nil {
		common.DisplayAppError(
			w,
			rErr,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	j, err := json.Marshal(questions)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}
