package data

import (
	"time"

	"github.com/timotew/etc/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type OptionRepository struct {
	C *mgo.Collection
}

func (r *OptionRepository) Create(option *models.Option) error {
	obj_id := bson.NewObjectId()
	option.ID = obj_id
	option.CreatedOn = time.Now()
	option.ModifiedOn = time.Now()
	option.Deleted = false
	err := r.C.Insert(&option)
	return err
}

func (r *OptionRepository) Update(option *models.Option) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": option.ID},
		bson.M{"$set": bson.M{
			"body":       option.Body,
			"modifiedOn": time.Now(),
		}})
	return err
}

func (r *OptionRepository) Remove(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *OptionRepository) Delete(id string) error {
	err := r.C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"deleted": true}})
	return err
}

func (r *OptionRepository) GetAll() (options []models.Option, err error) {
	err = r.C.Find(bson.M{"deleted": false}).All(&options)
	return
}

func (r *OptionRepository) GetByID(id string) (option models.Option, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&option)
	return
}

func (r *OptionRepository) GetByQuestion(questionID string) (questions []models.Option, err error) {
	err = r.C.Find(bson.M{"questionId": bson.ObjectIdHex(questionID)}).All(&questions)
	return
}
