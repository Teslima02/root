package routers

import (
	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/controllers"
	"github.com/urfave/negroni"
)

// TODO: CLASSES GET ID, DELET MANY CLASSES AND UPDATE MANY CLASSES
func SetPostRoutes(router *mux.Router) *mux.Router {
	classRouter := mux.NewRouter()
	classRouter.HandleFunc("/classes", controllers.CreateClass).Methods("POST")
	//classRouter.HandleFunc("/classes", controllers.CreateManyClass).Methods("POST")
	classRouter.HandleFunc("/classes/{id}", controllers.UpdateClass).Methods("PUT")
	classRouter.HandleFunc("/classes", controllers.GetClasses).Methods("GET")
	classRouter.HandleFunc("/classes/{id}", controllers.GetClassByID).Methods("GET")
	classRouter.HandleFunc("/classes/key/{key}", controllers.GetClassByKey).Methods("GET")
	classRouter.HandleFunc("/classes/master/{id}", controllers.GetClassesByMaster).Methods("GET")
	classRouter.HandleFunc("/classes/{id}", controllers.DeleteClass).Methods("DELETE")
	router.PathPrefix("/classes").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(classRouter),
	))
	return router
}
