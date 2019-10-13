package routers

import (
	"github.com/gorilla/mux"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/controllers"
	"github.com/urfave/negroni"
)

func SetQuestionRoutes(router *mux.Router) *mux.Router {
	questionRouter := mux.NewRouter()
	questionRouter.HandleFunc("/questions", controllers.CreateQuestion).Methods("POST")
	questionRouter.HandleFunc("/questions/{id}", controllers.UpdateQuestion).Methods("PUT")
	questionRouter.HandleFunc("/questions", controllers.GetQuestions).Methods("GET")
	questionRouter.HandleFunc("/questions/{id}", controllers.GetQuestionByID).Methods("GET")
	questionRouter.HandleFunc("/questions/tests/{id}", controllers.GetQuestionsByTest).Methods("GET")
	questionRouter.HandleFunc("/questions/{id}", controllers.DeleteQuestion).Methods("DELETE")
	router.PathPrefix("/questions").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(questionRouter),
	))
	return router
}
