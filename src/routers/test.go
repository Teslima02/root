package routers

import (
	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/controllers"
	"github.com/urfave/negroni"
)

func SetTestRoutes(router *mux.Router) *mux.Router {
	testRouter := mux.NewRouter()
	testRouter.HandleFunc("/tests", controllers.CreateTest).Methods("POST")
	testRouter.HandleFunc("/tests/{id}", controllers.UpdateTest).Methods("PUT")
	testRouter.HandleFunc("/tests", controllers.GetTests).Methods("GET")
	testRouter.HandleFunc("/tests/{id}", controllers.GetTestByID).Methods("GET")
	testRouter.HandleFunc("/tests/{id}/questions", controllers.GetTestQuestions).Methods("GET")
	testRouter.HandleFunc("/tests/{id}/session", controllers.GetTestSessionByTestID).Methods("GET")
	testRouter.HandleFunc("/tests/class/key/{key}", controllers.GetTestsByClassKey).Methods("GET")
	testRouter.HandleFunc("/tests/class/{id}", controllers.GetTestsByClass).Methods("GET")
	testRouter.HandleFunc("/tests/users/{id}", controllers.GetTestsByCreator).Methods("GET")
	testRouter.HandleFunc("/tests/{id}", controllers.DeleteTest).Methods("DELETE")
	router.PathPrefix("/tests").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(testRouter),
	))
	return router
}
