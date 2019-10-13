package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// bool, for JSON booleans
// float64, for JSON numbers
// string, for JSON strings
// []interface{}, for JSON arrays
// map[string]interface{}, for JSON objects
// nil for JSON null

type (
	/*User the structure for user type
	*This could be a master or a student
	 */
	User struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstName"`
		LastName     string        `json:"lastName"`
		Email        string        `json:"email"`
		Phone        string        `json:"phoneNumber"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `bson:"hashpassword,omitempty" json:"-"`
		Address      string        `json:"address,omitempty"`
		Username     string        `json:"username,omitempty"`
		DisplayPc    string        `json:"displaypic,omitempty" json:"displayPic"`
		Active       bool          `json:"active,omitempty"`
		Role         string        `json:"role,omitempty"`
		CreatedOn    time.Time     `json:"createdOn,omitempty"`
		Deleted      bool          `json:"deleted"`
		ModifiedOn   time.Time     `json:"modifiedOn,omitempty"`
	}

	/*Option is a choice available for
	*every question
	 */
	Option struct {
		ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
		QuestionID bson.ObjectId `bson:"questionId" json:"questionId"`
		Body       string        `json:"body"`
		Correct    bool          `json:"-"`
		Deleted    bool          `json:"deleted"`
		CreatedOn  time.Time     `json:"createdOn,omitempty"`
		ModifiedOn time.Time     `json:"modifiedOn,omitempty"`
	}

	/*Choice to monitor the choice a user
	*make s on an attempt to write test
	*
	 */
	Choice struct {
		ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OptionID   bson.ObjectId `bson:"optionid" json:"optionId"`
		UseID      bson.ObjectId `bson:"userid" json:"userId"`
		Correct    bool          `bson:"correct" json:"-"`
		Deleted    bool          `json:"deleted"`
		QuestionID bson.ObjectId `bson:"questionid" json:"questionId"`
		SessionID  bson.ObjectId `bson:"sessionid" json:"sessionId"`
		CreatedOn  time.Time     `json:"createdOn,omitempty"`
		ModifiedOn time.Time     `json:"modifiedOn,omitempty"`
	}

	/*Question struct for all question
	*Under every Test
	 */
	Question struct {
		ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TestID     bson.ObjectId `bson:"testid" json:"testId"`
		Body       string        `json:"body"`
		Deleted    bool          `json:"deleted"`
		CreatedOn  time.Time     `json:"createdOn,omitempty"`
		ModifiedOn time.Time     `json:"modifiedOn,omitempty"`
		Point      int           `json:"point"`
	}

	/*Test is the assessment that is to be done bu student
	*created by the master
	 */
	Test struct {
		ID                     bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy              bson.ObjectId `bson:"createdBy"`
		Title                  string        `json:"title"`
		Description            string        `json:"description"`
		QuestionsPerAssessment int           `json:"questionsPerAssessment"`
		ClassID                bson.ObjectId `bson:"classId" json:"classId"`
		CreatedOn              time.Time     `json:"createdOn,omitempty"`
		AccessDate             time.Time     `json:"accessDate,omitempty"`
		DueDate                time.Time     `json:"dueDate,omitempty"`
		Deleted                bool          `json:"deleted"`
		ModifiedOn             time.Time     `json:"modifiedOn,omitempty"` // 02-11-2017 15:04:05
		Duration               time.Duration `json:"duration,omitempty"`   // 9m
		Status                 string        `json:"status,omitempty"`
		Tags                   []string      `json:"tags,omitempty"`
	}

	// TestSession to trace the test session and time
	TestSession struct {
		ID         bson.ObjectId   `bson:"_id,omitempty" json:"id"`
		TestID     bson.ObjectId   `bson:"testId,omitempty" json:"testId"`
		UserID     bson.ObjectId   `bson:"userId,omitempty" json:"userId"`
		Questions  []bson.ObjectId `bson:"questions,omitempty" json:"-"`
		Time       time.Duration   `json:"time,omitempty"`
		Active     bool            `json:"active"`
		Deleted    bool            `json:"deleted"`
		CreatedOn  time.Time       `json:"createdOn,omitempty"`
		ModifiedOn time.Time       `json:"modifiedOn,omitempty"`
	}

	/*Class is the struct that keeps every student is a class
	*and give them access to some assessments
	 */
	Class struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		MasterID    bson.ObjectId `bson:"masterId" json:"masterId"`
		ClassName   string        `json:"className"`
		Description string        `json:"description"`
		Alias       string        `json:"code"` //i.e MEC 103
		Key         string        `json:"key"`
		Published   bool          `json:"published"`
		Deleted     bool          `json:"deleted"`
		CreatedOn   time.Time     `json:"createdOn,omitempty"`
		ModifiedOn  time.Time     `json:"modifiedOn,omitempty"`
	}

	/*Result are the  every test
	*It the core
	 */
	Result struct {
		ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
		SessionID  bson.ObjectId `bson:"sessionId" json:"sessionId"`
		Points     int           `json:"points"`
		Status     string        `json:"status"` //held, release e.t.c
		Deleted    bool          `json:"deleted"`
		CreatedOn  time.Time     `json:"createdOn,omitempty"`
		ModifiedOn time.Time     `json:"modifiedOn,omitempty"`
	}
)
