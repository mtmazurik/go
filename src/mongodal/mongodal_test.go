package mongodal_test 

import (
	"testing"
	"time"
	"fmt"
	"encoding/json"

	"mongodal"
	"typelib"
)
// TestCreate tests inserting a record in Mongo Atlas
func TestCreate(t *testing.T){
	
	repo := LoadRepoObject() // static data to test with

	s,_ := json.Marshal(repo)
	fmt.Printf("%s",string(s))
	
	insertedID, err := mongodal.Create(repo)
	if  err!= nil {
		fmt.Printf("Test failed.")
	}

	fmt.Println("Inserted a single document: ", insertedID)
}
// LoadRepoObject is an object loader
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
