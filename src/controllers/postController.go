package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")
}

func allArticles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "GET" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	}
	//  else {
	// Your code goes here
	// }

	fmt.Println("Endpoint: All article endpoint")
	json.NewEncoder(w).Encode(Articles)
}

func oneArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// fmt.Fprint(w, "key: "+key)

	// Loop over all of our Articles
	// if the article.ID equals the key we pass in
	// return the article encoded as JSON

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {

	// fmt.Print(r.Body)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "POST" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	}
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
	// fmt.Fprintf(w, "%+v", string(reqBody))
}

func updateArticle(w http.ResponseWriter, r *http.Request) {

	fmt.Print(r)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "PUT" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	}

	vars := mux.Vars(r)
	key := vars["id"]

	// json.NewEncoder(w).Encode(article)
	fmt.Fprint(w, "key: "+key)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)

	// we will need to extract the `id` of the article we
	// wish to delete

	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if article.ID == id {
			// updates our Articles array to remove the
			// article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}
