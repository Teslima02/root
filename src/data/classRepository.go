package data

import (
	"time"

	"github.com/timotew/etc/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ClassRepository struct {
	C *mgo.Collection
}

func (r *ClassRepository) Create(class *models.Class) error {
	obj_id := bson.NewObjectId()
	class.ID = obj_id
	class.CreatedOn = time.Now()
	class.ModifiedOn = time.Now()
	class.Deleted = false
	err := r.C.Insert(&class)
	return err
}

func (r *ClassRepository) CreateMany(classes []*models.Class) error {
	new := make([]interface{}, len(classes))
	for i, v := range classes {
		new[i] = v
	}
	err := r.C.Insert(new...)
	return err
}

func (r *ClassRepository) Update(class *models.Class) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": class.ID},
		bson.M{"$set": bson.M{
			"classname":  class.ClassName,
			"alias":      class.Alias,
			"masterId":   class.MasterID,
			"modifiedOn": time.Now(),
		}})
	return err
}
func (r *ClassRepository) Remove(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *ClassRepository) Delete(id string) error {
	err := r.C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"deleted": true}})
	return err
}

func (r *ClassRepository) GetAll() (classes []models.Class, err error) {
	err = r.C.Find(bson.M{"deleted": false}).All(&classes)
	return
}

func (r *ClassRepository) GetByAlias(alias string) (classes []models.Class, err error) {
	err = r.C.Find(bson.M{"alias": alias}).All(&classes)
	return
}

func (r *ClassRepository) GetById(id string) (class models.Class, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&class)
	return
}

func (r *ClassRepository) GetByKey(key string) (class models.Class, err error) {
	err = r.C.Find(bson.M{"key": key, "deleted": false}).One(&class)
	return
}

func (r *ClassRepository) GetByMaster(master string) (classes []models.Class, err error) {
	err = r.C.Find(bson.M{"masterId": bson.ObjectIdHex(master), "deleted": false}).All(&classes)
	return
}
