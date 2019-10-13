package routers

import (
	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/controllers"
	"github.com/urfave/negroni"
)

// SetUserRoutes route for user controller
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	proRouter := mux.NewRouter()
	proRouter.HandleFunc("/users/verify", controllers.Verify).Methods("GET")
	proRouter.HandleFunc("/users/tmpup", controllers.TempUserUpdate).Methods("POST")
	proRouter.HandleFunc("/users/find", controllers.Find).Methods("POST")
	proRouter.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	router.PathPrefix("/users/verify").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(proRouter),
	))
	router.PathPrefix("/users/find").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(proRouter),
	))
	router.PathPrefix("/users/{id}").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(proRouter),
	))
	return router
}
