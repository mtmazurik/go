package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository data object for storage in MongoDB (Atlas-hosted solution)
type Repository struct {
	ID           int
	KeyName      string
	KeyValue     string
	Tags         [10]string
	CreatedDate  time.Time
	CreatedBy    string
	ModifiedDate time.Time
	ModifiedBy   string
	App          string
	Repository   string
	Collection   string
	SchemaName   string
	SchemaURI    string
	Body         string
}

func main() {

	clientOptions := options.Client().ApplyURI("mongo-db-connection-string-goes-here-occluded-for-security")

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
	repo.SchemaURI = "https://www.cloudcomputingassociates.com/schemas/repository.schema.json"
	repo.Body = "{ \"flight\" : { \"flightNumber\" : \"580\", \"dayOfWeek\" : \"Monday\",\"manufacturer\":\"Boeing\", \"planeType\":\"777x\" }}"

	return repo
}
