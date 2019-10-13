package routers

import (
	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/controllers"
	"github.com/timotew/etc/src/data"
	"github.com/urfave/negroni"
)

func SetTestSessionRoutes(router *mux.Router) *mux.Router {

	testRoom := &data.SocketRoom{
		Forward: make(chan []byte),
		Join:    make(chan *data.SocketClient),
		Leave:   make(chan *data.SocketClient),
		Clients: make(map[*data.SocketClient]bool),
	}
	router.Handle("/test_sessions_pool/{user}", testRoom)
	// get the room going
	go testRoom.Run()
	testSessionRouter := mux.NewRouter()
	testSessionRouter.HandleFunc("/test_sessions", controllers.CreateTestSession).Methods("POST")
	testSessionRouter.HandleFunc("/test_sessions/{id}", controllers.UpdateTestSession).Methods("PUT")
	testSessionRouter.HandleFunc("/test_sessions/{id}", controllers.GetTestSessionByID).Methods("GET")
	testSessionRouter.HandleFunc("/test_sessions/{id}/questions", controllers.GetTestSessionQuestions).Methods("GET")
	testSessionRouter.HandleFunc("/test_sessions/{id}/close", controllers.CloseTestSession).Methods("PUT")
	router.PathPrefix("/test_sessions").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(testSessionRouter),
	))
	return router
}
