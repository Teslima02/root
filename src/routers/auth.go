package routers

import (
	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/controllers"
)

// SetUserRoutes route for user controller
func SetAuthRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/kit/partial", controllers.FBKitPartial).Methods("POST")
	router.HandleFunc("/auth/kit/login", controllers.FBKitLogin).Methods("POST")
	// proRouter := mux.NewRouter()
	// proRouter.HandleFunc("/users/verify", controllers.Verify).Methods("GET")
	// proRouter.HandleFunc("/users/find", controllers.Find).Methods("POST")
	// proRouter.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	// router.PathPrefix("/users/verify").Handler(negroni.New(
	// 	negroni.HandlerFunc(common.Authorize),
	// 	negroni.Wrap(proRouter),
	// ))
	// router.PathPrefix("/users/find").Handler(negroni.New(
	// 	negroni.HandlerFunc(common.Authorize),
	// 	negroni.Wrap(proRouter),
	// ))
	// router.PathPrefix("/users/{id}").Handler(negroni.New(
	// 	negroni.HandlerFunc(common.Authorize),
	// 	negroni.Wrap(proRouter),
	// ))
	return router
}
