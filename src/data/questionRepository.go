package data

import (
	"time"

	"github.com/timotew/etc/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type QuestionRepository struct {
	C *mgo.Collection
}

func (r *QuestionRepository) Create(question *models.Question) error {
	objID := bson.NewObjectId()
	question.ID = objID
	question.CreatedOn = time.Now()
	question.ModifiedOn = time.Now()
	question.Deleted = false
	err := r.C.Insert(&question)
	return err
}

func (r *QuestionRepository) Update(question *models.Question) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": question.ID},
		bson.M{"$set": bson.M{
			"body":       question.Body,
			"point":      question.Point,
			"modifiedOn": time.Now(),
		}})
	return err
}

func (r *QuestionRepository) Remove(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *QuestionRepository) Delete(id string) error {
	err := r.C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"deleted": true}})
	return err
}

func (r *QuestionRepository) GetManyByObjectId(ids []bson.ObjectId) (questions []models.Question, err error) {
	err = r.C.Find(bson.M{"_id": bson.M{"$in": ids}}).All(&questions)
	return
}

func (r *QuestionRepository) GetAll() (questions []models.Question, err error) {
	err = r.C.Find(bson.M{"deleted": false}).All(&questions)
	return
}

func (r *QuestionRepository) GetById(id string) (question models.Question, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&question)
	return
}

func (r *QuestionRepository) GetByTest(testID string) []models.Question {
	var questions []models.Question
	iter := r.C.Find(bson.M{"testId": bson.ObjectIdHex(testID)}).Iter()
	result := models.Question{}
	for iter.Next(&result) {
		questions = append(questions, result)
	}
	return questions
}

func (r *QuestionRepository) GetRandomQuestionsIDsByTestID(testID string, limit int) (questionIds []models.Question, err error) {
	err = r.C.Find(bson.M{"testId": bson.ObjectIdHex(testID)}).Limit(limit).Sort("$natural").Select(bson.M{"_id": 1}).All(&questionIds)
	return
}
