package data

import (
	"time"

	"github.com/timotew/etc/src/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	C *mgo.Collection
}

func (r *UserRepository) UpdateTmpUser(user *models.User) error {

	err := r.C.Update(bson.M{"_id": user.ID},
		bson.M{"$set": bson.M{
			"firstname":  user.FirstName,
			"username":   user.Username,
			"modifiedon": time.Now(),
		}})
	return err
}

func (r *UserRepository) CreateUser(user *models.User) error {
	objID := bson.NewObjectId()
	user.ID = objID
	user.Deleted = false
	user.CreatedOn = time.Now()
	user.ModifiedOn = time.Now()
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	// clear the incoming text password
	user.Password = ""
	err = r.C.Insert(&user)
	return err
}

func (r *UserRepository) SearchUserByPhone(phone string) (u models.User, err error) {
	err = r.C.Find(bson.M{"phoneNumber": phone}).One(&u)
	return
}

func (r *UserRepository) SearchByUsername(text string, inList []string) (u []*models.User, err error) {
	err = r.C.Find(
		bson.M{"$and": []bson.M{
			bson.M{"email": bson.M{"$regex": bson.RegEx{text, ""}}}, bson.M{"email": bson.M{"$nin": inList}}}}).Select(bson.M{"email": 1}).All(&u)
	if err != nil {
		return
	}
	return
}

func (r *UserRepository) GetByID(ID string) (u models.User, err error) {
	err = r.C.FindId(bson.ObjectIdHex(ID)).One(&u)
	return
}

func (r *UserRepository) Delete(id string) error {
	err := r.C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"deleted": true}})
	return err
}

func (r *UserRepository) Remove(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *UserRepository) RemoveAll() error {
	_, err := r.C.RemoveAll(bson.M{})
	return err
}

func (r *UserRepository) Login(user models.User) (u models.User, err error) {

	err = r.C.Find(bson.M{"email": user.Email, "deleted": false}).One(&u)
	if err != nil {
		return
	}
	// Validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
	}
	return
}
