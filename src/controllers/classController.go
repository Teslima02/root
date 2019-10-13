package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/data"
	"github.com/timotew/etc/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//CreateClass Handler for HTTP Post - "/class"
// Insert a new Note document for a TaskId
func CreateClass(w http.ResponseWriter, r *http.Request) {
	//get class resource
	var dataResource ClassResource
	// Decode the incoming tEST json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Class data", 500)
		return
	}
	classModel := dataResource.Data
	context := NewContext()
	defer context.Close()
	ctx := r.Context()
	if userID := ctx.Value("user"); userID != nil {
		context.User = userID.(string)
	} else {
		common.DisplayAppError(
			w,
			errors.New("Invalid User ID"),
			"Error unrecognized context",
			500,
		)
		return
	}
	class := &models.Class{
		MasterID:    bson.ObjectIdHex(context.User),
		ClassName:   classModel.ClassName,
		Alias:       classModel.Alias,
		Description: classModel.Description,
		Published:   classModel.Published,
	}
	col := context.DbCollection("classes")
	//Insert a class document
	repo := &data.ClassRepository{C: col}

	// check for class duplicate
	classes, _ := repo.GetByAlias(classModel.Alias)
	// check if owner has duplicate
	for _, _class := range classes {
		if _class.MasterID == class.MasterID {
			common.DisplayAppError(w, errors.New("Class already exists"), "Class already Exists.", 409)
			return
		}
	}

	if len(classes) > 0 {
		num := len(classes)
		class.Key = strings.Join([]string{class.Alias, strconv.Itoa(num)}, "-")
	} else {
		class.Key = class.Alias
	}

	iErr := repo.Create(class)
	if iErr != nil {
		common.DisplayAppError(
			w,
			iErr,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	j, err := json.Marshal(class)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func CreateManyClass(w http.ResponseWriter, r *http.Request) {
	//get class resource
	var dataResource ManyClassResource
	// Decode the incoming tEST json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Class data", 500)
		return
	}

	var classModels []*models.Class
	for _, classModel := range dataResource.Data {
		class := &models.Class{
			ID:         bson.NewObjectId(),
			MasterID:   bson.ObjectIdHex(classModel.MasterID),
			ClassName:  classModel.ClassName,
			Alias:      classModel.Alias,
			CreatedOn:  time.Now(),
			ModifiedOn: time.Now(),
		}

		classModels = append(classModels, class)
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("classes")
	repo := &data.ClassRepository{C: col}
	iErr := repo.CreateMany(classModels)
	if iErr != nil {
		common.DisplayAppError(
			w,
			iErr,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	j, err := json.Marshal(classModels)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//GetClassesByMaster Handler for HTTP Get - "/classes/users/{id}
// Returns all tests documents under a QuestionID
func GetClassesByMaster(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	isHex := bson.IsObjectIdHex(id)
	if !isHex {
		common.DisplayAppError(
			w,
			errors.New("Invalid Record ID"),
			"Record Not found",
			http.StatusNotFound,
		)
		return
	}
	col := context.DbCollection("classes")
	repo := &data.ClassRepository{C: col}
	classes, _ := repo.GetByMaster(id)
	j, err := json.Marshal(ClassesResource{Data: classes})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetClasses Handler for HTTP Get - "/classes"
// Returns all classes documents
func GetClasses(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("classes")
	repo := &data.ClassRepository{C: col}
	classes, _ := repo.GetAll()
	j, err := json.Marshal(ClassesResource{Data: classes})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetClassByID Handler for HTTP Get - "/classes/{id}"
// Returns a single Class document by id
func GetClassByID(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("classes")
	repo := &data.ClassRepository{C: col}
	isHex := bson.IsObjectIdHex(id)
	if !isHex {
		common.DisplayAppError(
			w,
			errors.New("Invalid Record ID"),
			"Record Not found",
			http.StatusNotFound,
		)
		return
	}
	class, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return

	}

	j, err := json.Marshal(class)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

func GetClassByKey(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	key := vars["key"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("classes")
	repo := &data.ClassRepository{C: col}
	class, err := repo.GetByKey(key)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return

	}

	j, err := json.Marshal(class)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//UpdateClass Handler for HTTP Put - "/classes/{id}"
// Update an existing Class document
func UpdateClass(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	isHex := bson.IsObjectIdHex(vars["id"])
	if !isHex {
		common.DisplayAppError(
			w,
			errors.New("Invalid Record ID"),
			"Record Not found",
			http.StatusNotFound,
		)
		return
	}
	id := bson.ObjectIdHex(vars["id"])
	var dataResource ClassResource
	// Decode the incoming Class json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Class data", 500)
		return
	}
	classModel := dataResource.Data
	class := &models.Class{
		ID:        id,
		MasterID:  bson.ObjectIdHex(classModel.MasterID),
		ClassName: classModel.ClassName,
		Alias:     classModel.Alias,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("classes")
	repo := &data.ClassRepository{C: col}
	//Update test document
	if err := repo.Update(class); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

//DeleteClass Handler for HTTP Delete - "/classes/{id}"
// Delete an existing Class document
func DeleteClass(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	isHex := bson.IsObjectIdHex(id)
	if !isHex {
		common.DisplayAppError(
			w,
			errors.New("Invalid Record ID"),
			"Record Not found",
			http.StatusNotFound,
		)
		return
	}
	col := context.DbCollection("classes")
	repo := &data.ClassRepository{C: col}
	//Delete a class document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
