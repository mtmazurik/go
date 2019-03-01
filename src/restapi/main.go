package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	log.Fatal(http.ListenAndServe(":8901", router))
}

// Person : is for person
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address : is for addresses
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// GetPeople : gets people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	people = append(people, Person{ID: "1", Firstname: "Marty", Lastname: "Mazurik", Address: &Address{City: "Happy Valley", State: "OR"}})
	json.NewEncoder(w).Encode(people)
}
