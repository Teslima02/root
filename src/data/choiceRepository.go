package data

import (
	"time"

	"github.com/timotew/etc/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ChoiceRepository struct {
	C *mgo.Collection
}

func (r *ChoiceRepository) Create(choice *models.Choice) error {
	obj_id := bson.NewObjectId()
	choice.ID = obj_id
	choice.CreatedOn = time.Now()
	choice.ModifiedOn = time.Now()
	choice.Deleted = false
	err := r.C.Insert(&choice)
	return err
}

func (r *ChoiceRepository) Update(choice *models.Choice) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": choice.ID},
		bson.M{"$set": bson.M{
			"correct":    choice.Correct,
			"modifiedon": time.Now(),
		}})
	return err
}

func (r *ChoiceRepository) Delete(id string) error {
	err := r.C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"deleted": true}})
	return err
}

func (r *ChoiceRepository) Remove(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *ChoiceRepository) GetAll() (choices []models.Choice, err error) {
	err = r.C.Find(bson.M{"deleted": false}).All(&choices)
	return
}

func (r *ChoiceRepository) GetById(id string) (choice models.Choice, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&choice)
	return
}

func (r *ChoiceRepository) GetByQuestion(questionID string) []models.Choice {
	var choices []models.Choice
	iter := r.C.Find(bson.M{"questionid": questionID}).Iter()
	result := models.Choice{}
	for iter.Next(&result) {
		choices = append(choices, result)
	}
	return choices
}

func (r *ChoiceRepository) GetByQuestionID(questionID string, userID string) (choice models.Choice, err error) {
	err = r.C.Find(bson.M{"questionId": bson.ObjectIdHex(questionID), "userId": bson.ObjectIdHex(userID)}).One(&choice)
	return
}

func (r *ChoiceRepository) GetByQuestionOption(questionID, optionID string) (choice models.Choice, err error) {
	err = r.C.Find(bson.M{"questionId": bson.ObjectIdHex(questionID), "optionId": bson.ObjectIdHex(optionID)}).One(&choice)
	return
}
