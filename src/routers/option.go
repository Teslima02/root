package routers

import (
	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/controllers"
	"github.com/urfave/negroni"
)

func SetOptionRoutes(router *mux.Router) *mux.Router {
	optionRouter := mux.NewRouter()
	optionRouter.HandleFunc("/options", controllers.CreateOption).Methods("POST")
	optionRouter.HandleFunc("/options/{id}", controllers.UpdateOption).Methods("PUT")
	optionRouter.HandleFunc("/options", controllers.GetOptions).Methods("GET")
	optionRouter.HandleFunc("/options/{id}", controllers.GetOptionByID).Methods("GET")
	optionRouter.HandleFunc("/options/questions/{id}", controllers.GetOptionsByQuestion).Methods("GET")
	optionRouter.HandleFunc("/options/{id}", controllers.DeleteOption).Methods("DELETE")
	router.PathPrefix("/options").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(optionRouter),
	))
	return router
}
