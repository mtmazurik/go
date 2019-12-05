package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"

	"mongodal"
	"typelib"
)

// main is entry point to API - using Gorilla Mux, see more here https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", create ).Methods("POST")
	log.Fatal(http.ListenAndServe(":8901", router))
}
func create(w http.ResponseWriter, r *http.Request) {
	var repo typelib.Repository
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &repo)

	insertedID, err := mongodal.Create(repo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertedID)
	

}

