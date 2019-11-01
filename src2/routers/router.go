package routers

import (
	"github.com/gorilla/mux"
)

// InitRoutes initialize all routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the Class entity
	router = SetPostRoutes(router)
	return router
}
