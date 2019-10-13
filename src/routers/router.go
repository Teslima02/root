package routers

import (
	"github.com/gorilla/mux"
)

// InitRoutes initialize all routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for the Test entity
	router = SetTestRoutes(router)
	// Routes for the Question entity
	router = SetQuestionRoutes(router)
	// Routes for the Option entity
	router = SetOptionRoutes(router)
	// Routes for the Choice entity
	router = SetChoiceRoutes(router)
	// Routes for the Result entity
	router = SetResultRoutes(router)
	// Routes for the Result entity
	router = SetClassRoutes(router)
	// Routes for testSession
	router = SetTestSessionRoutes(router)
	// AUTH KITS
	router = SetAuthRoutes(router)
	return router
}
