package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xeipuuv/gojsonschema"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository data object for storage in MongoDB (Atlas-hosted solution)
type Repository struct {
	ID           int        `json:"_id"`
	KeyName      string     `json:"keyName"`
	KeyValue     string     `json:"keyValue"`
	Tags         [10]string `json:"tags"`
	CreatedDate  time.Time  `json:"createdDate"`
	CreatedBy    string     `json:"createdBy"`
	ModifiedDate time.Time  `json:"modifiedDate"`
	ModifiedBy   string     `json:"modifiedBy"`
	App          string     `json:"app"`
	Repository   string     `json:"repository"`
	Collection   string     `json:"collection"`
	SchemaName   string     `json:"schemaName"`
	SchemaURI    string     `json:"schemaURI"`
	Body         string     `json:"body"`
}

func main() {

	clientOptions := options.Client().ApplyURI("MONGO ATLAS connection string occluded for safety")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	//Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	//Access the collection

	collection := client.Database("flight-db").Collection("flights")

	repo := LoadRepoObject()
	r, err := json.Marshal(repo)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	s := string(r)

	err = ValidateJSON(&s, &repo.SchemaURI)
	if err != nil {
		os.Exit(3) // alternative to return
	}

	insertResult, err := collection.InsertOne(context.TODO(), repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// you'll want to use pooling, but for now just disconnect the client on this one-pass code example
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

}

// ValidateJSON uses JSON Schema by xeipusanthosh-tekuri https://github.com/santhosh-tekuri/jsonschema
func ValidateJSON(s *string, schemaURI *string) error {

	documentLoader := gojsonschema.NewStringLoader(*s)
	schemaLoader := gojsonschema.NewReferenceLoader(*schemaURI)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is invalid. see errors:\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

	return nil
}

// LoadRepoObject is an object loader
func LoadRepoObject() Repository {

	var repo Repository

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
