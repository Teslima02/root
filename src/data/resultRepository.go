package data

import (
	"time"

	"github.com/timotew/etc/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ResultRepository struct {
	C *mgo.Collection
}

func (r *ResultRepository) Create(result *models.Result) error {
	obj_id := bson.NewObjectId()
	result.ID = obj_id
	result.CreatedOn = time.Now()
	result.ModifiedOn = time.Now()
	result.Deleted = false
	err := r.C.Insert(&result)
	return err
}

func (r *ResultRepository) Update(result *models.Result) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": result.ID},
		bson.M{"$set": bson.M{
			"points":     result.Points,
			"status":     result.Status,
			"modifiedOn": time.Now(),
		}})
	return err
}

func (r *ResultRepository) Remove(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *ResultRepository) Delete(id string) error {
	err := r.C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"deleted": true}})
	return err
}

func (r *ResultRepository) GetAll() (results []models.Result, err error) {
	err = r.C.Find(bson.M{"deleted": false}).All(&results)
	return
}

func (r *ResultRepository) GetById(id string) (result models.Result, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (r *ResultRepository) GetByTest(testID string) []models.Result {
	var results []models.Result
	iter := r.C.Find(bson.M{"testId": testID}).Iter()
	result := models.Result{}
	for iter.Next(&result) {
		results = append(results, result)
	}
	return results
}
