package common

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// GetSession get current database session
func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoDBHost},
			Username: AppConfig.DBUser,
			Password: AppConfig.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}
func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}

// Add indexes into MongoDB
func addIndexes() {
	var err error
	userIndex := mgo.Index{
		Key:        []string{"email", "username"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	testIndex := mgo.Index{
		Key:        []string{"createdBy", "classId"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}
	questionIndex := mgo.Index{
		Key:        []string{"testId"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}
	optionIndex := mgo.Index{
		Key:        []string{"questionId"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}

	choiceIndex := mgo.Index{
		Key:        []string{"questionId", "optionId", "sessionId"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}

	resultIndex := mgo.Index{
		Key:        []string{"testId", "userId"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}

	classIndex := mgo.Index{
		Key:        []string{"masterId"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}

	testSessionIndex := mgo.Index{
		Key:        []string{"userId", "testId"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}

	// Add indexes into MongoDB
	session := GetSession().Copy()
	defer session.Close()
	userCol := session.DB(AppConfig.Database).C("users")
	testCol := session.DB(AppConfig.Database).C("tests")
	questionCol := session.DB(AppConfig.Database).C("questions")
	optionCol := session.DB(AppConfig.Database).C("options")
	choiceCol := session.DB(AppConfig.Database).C("choices")
	resultCol := session.DB(AppConfig.Database).C("results")
	classCol := session.DB(AppConfig.Database).C("classes")
	testSessionCol := session.DB(AppConfig.Database).C("testSessions")

	err = userCol.EnsureIndex(userIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = testCol.EnsureIndex(testIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = questionCol.EnsureIndex(questionIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = optionCol.EnsureIndex(optionIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = choiceCol.EnsureIndex(choiceIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = resultCol.EnsureIndex(resultIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = classCol.EnsureIndex(classIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = testSessionCol.EnsureIndex(testSessionIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
}
