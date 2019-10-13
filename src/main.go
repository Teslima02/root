package main

import (
	"log"
	"net/http"

	common "github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/routers"
	"github.com/urfave/negroni"
)

//Entry point of the program
func main() {

	//common.StartUp() - Replaced with init method
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}
