package routers

import (
	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/controllers"
	"github.com/urfave/negroni"
)

func SetChoiceRoutes(router *mux.Router) *mux.Router {
	choiceRouter := mux.NewRouter()
	choiceRouter.HandleFunc("/choices", controllers.CreateChoice).Methods("POST")
	choiceRouter.HandleFunc("/choices/{id}", controllers.UpdateChoice).Methods("PUT")
	choiceRouter.HandleFunc("/choices", controllers.GetChoices).Methods("GET")
	choiceRouter.HandleFunc("/choices/{id}", controllers.GetChoiceByID).Methods("GET")
	choiceRouter.HandleFunc("/choices/questions/{id}", controllers.GetChoicesByQuestion).Methods("GET")
	choiceRouter.HandleFunc("/choices/{id}", controllers.DeleteChoice).Methods("DELETE")
	router.PathPrefix("/choices").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(choiceRouter),
	))
	return router
}
