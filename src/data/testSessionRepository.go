package data

import (
	"time"

	"github.com/timotew/etc/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TestSessionRepository struct {
	C *mgo.Collection
}

func (r *TestSessionRepository) Create(session *models.TestSession) error {
	objID := bson.NewObjectId()
	session.ID = objID
	session.CreatedOn = time.Now()
	session.ModifiedOn = time.Now()
	session.Deleted = false
	session.Active = true
	err := r.C.Insert(&session)
	return err
}

func (r *TestSessionRepository) Update(session *models.TestSession) error {
	err := r.C.Update(bson.M{"_id": session.ID},
		bson.M{"$set": bson.M{
			"time":       session.Time,
			"active":     session.Active,
			"modifiedon": time.Now(),
		}})
	return err
}

func (r *TestSessionRepository) GetById(id string) (test models.TestSession, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&test)
	return
}

func (r *TestSessionRepository) GetByTestId(testID, userID string) (testSession models.TestSession, err error) {
	err = r.C.Find(bson.M{"testId": bson.ObjectIdHex(testID), "userId": bson.ObjectIdHex(userID)}).One(&testSession)
	return
}

func (r *TestSessionRepository) CloseSession(testID, userID string) error {
	err := r.C.Update(bson.M{"_id": bson.ObjectIdHex(testID), "userId": bson.ObjectIdHex(userID)},
		bson.M{"$set": bson.M{
			"active":     false,
			"modifiedon": time.Now(),
		}})
	return err
}
