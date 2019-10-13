package routers

import (
	"github.com/gorilla/mux"
	"github.com/timotew/beacon/common"
	"github.com/timotew/beacon/controllers"
	"github.com/urfave/negroni"
)

func SetResultRoutes(router *mux.Router) *mux.Router {
	resultRouter := mux.NewRouter()
	resultRouter.HandleFunc("/results", controllers.CreateResult).Methods("POST")
	resultRouter.HandleFunc("/results/{id}", controllers.UpdateResult).Methods("PUT")
	resultRouter.HandleFunc("/results", controllers.GetResults).Methods("GET")
	resultRouter.HandleFunc("/results/{id}", controllers.GetResultByID).Methods("GET")
	resultRouter.HandleFunc("/results/tests/{id}", controllers.GetResultsByTest).Methods("GET")
	resultRouter.HandleFunc("/results/{id}", controllers.DeleteResult).Methods("DELETE")
	router.PathPrefix("/results").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(resultRouter),
	))
	return router
}
