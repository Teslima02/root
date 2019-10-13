package data

import (
	"time"

	"github.com/timotew/etc/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TestRepository struct {
	C *mgo.Collection
}

func (r *TestRepository) Create(test *models.Test) error {
	objID := bson.NewObjectId()
	test.ID = objID
	test.CreatedOn = time.Now()
	test.ModifiedOn = time.Now()
	test.Deleted = false
	err := r.C.Insert(&test)
	return err
}

func (r *TestRepository) Update(test *models.Test) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": test.ID},
		bson.M{"$set": bson.M{
			"title":                  test.Title,
			"description":            test.Description,
			"classId":                test.ClassID,
			"questionsPerAssessment": test.QuestionsPerAssessment,
			"accessDate":             test.AccessDate,
			"dueDate":                test.DueDate,
			"duration":               test.Duration,
			"status":                 test.Status,
			"tags":                   test.Tags,
			"modifiedOn":             time.Now(),
		}})
	return err
}

func (r *TestRepository) Remove(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *TestRepository) Delete(id string) error {
	err := r.C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"deleted": true}})
	return err
}

func (r *TestRepository) GetByClass(id string) (tests []models.Test, err error) {
	err = r.C.Find(bson.M{"classId": bson.ObjectIdHex(id)}).All(&tests)
	return
}

func (r *TestRepository) GetByCreator(id string) (tests []models.Test, err error) {
	err = r.C.Find(bson.M{"deleted": false, "createdby": bson.ObjectIdHex(id)}).All(&tests)
	return
}

/*GetAll test available
 */
func (r *TestRepository) GetAll() (tests []models.Test, err error) {
	err = r.C.Find(bson.M{"deleted": false}).All(&tests)
	return
}

func (r *TestRepository) GetByClassKey(key string) (tests []models.Test, err error) {
	err = r.C.Find(bson.M{"key": key, "deleted": false}).All(&tests)
	return
}

func (r *TestRepository) GetById(id string) (test models.Test, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&test)
	return
}
