package mongodal_test 

import (
	"testing"
	"time"
	"fmt"
	//"encoding/json"

	"mongodal"
	"typelib"
)
// TestCreate tests insert record into Mongo:repository-nook-db / test-collection
func TestCreate(t *testing.T){
	
	cn := "***** occluded for safety:  mongo connection string to Atlas (hosted) *****"
	repo := LoadRepoObject() // static data to test with
	
	//s,_ := json.Marshal(repo)
	//fmt.Printf("%s",string(s))
	
	insertedID, err := mongodal.Create(cn, "repository-nook-db", "test-collection", repo)
	if  err!= nil {
		fmt.Printf("Test failed.")
	}

	fmt.Println("Inserted a single document: ", insertedID)
}
// LoadRepoObject inserts dummy data into struct
func LoadRepoObject() typelib.Repository {

	var repo typelib.Repository

	repo.ID = 0
	repo.KeyName = "flightNumber"
	repo.KeyValue = "999"
	repo.Tags[0] = "planeType:777x"
	repo.Tags[1] = "manufacturer:Boeing"
	repo.CreatedDate = time.Now()
	repo.CreatedBy = "Marty Mazurik"
	repo.ModifiedDate = time.Now()
	repo.ModifiedBy = "Marty Mazurik"
	repo.App = "RepositoryNook"
	repo.Repository = "flight-db"
	repo.Collection = "flights"
	repo.SchemaName = "flight"
	repo.SchemaURI = "http://www.cloudcomputingassociates.com/schemas/repository.schema.json"
	repo.Body = "{ \"flight\" : { \"flightNumber\" : \"580\", \"dayOfWeek\" : \"Monday\",\"manufacturer\":\"Boeing\", \"planeType\":\"777x\" }}"

	return repo
}
