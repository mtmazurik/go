package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"

	"mongodal"		// mongo Data Access Layer package
	"typelib"
)

var cn = "***** occluded for safety:  mongo connection string to Atlas (hosted) *****"

// main is entry point to API - using Gorilla Mux, see more here https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{db}/{coll}", create ).Methods("POST")
	log.Fatal(http.ListenAndServe(":8901", router))
}
func create(w http.ResponseWriter, r *http.Request) {
	resources := mux.Vars(r)
	var repo typelib.Repository
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &repo)
	db := resources["db"]
	coll := resources["coll"]
	insertedID, err := mongodal.Create(cn,db,coll,repo)		// insert doc into Mongo

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertedID)
	
}

