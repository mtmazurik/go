package mongodal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/xeipuuv/gojsonschema"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"typelib"
)

// Create - Insert document; see MongoDB GitHub readme for more info on Driver usage https://github.com/mongodb/mongo-go-driver#usage 
func Create(repo typelib.Repository) (string, error) {

	clientOptions := options.Client().ApplyURI("occluded for safety-mongodb atlas connection string here")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil) 	//Check the connection

	if err != nil {
		log.Fatal(err)
	} 

	collection := client.Database("flight-db").Collection("flights")

	r, err := json.Marshal(repo)
	if err != nil {
		fmt.Printf("Error: %s", err)
		log.Fatal(err)
	}
	s := string(r)

	err = ValidateJSON(&s, &repo.SchemaURI)
	if err != nil {
		log.Fatal(err)
	}

	insertResult, err := collection.InsertOne(context.TODO(), repo)
	if err != nil {
		log.Fatal(err)
	}

	// you'll want to use pooling, but for now just disconnect the client on this one-pass code example
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
// convert the returned Object from Mongo, see more in this post: https://stackoverflow.com/questions/49933249/mongo-go-driver-get-objectid-from-insert-result
    var returnIDstring string
	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		returnIDstring = oid.Hex()
	}
	return returnIDstring, nil
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
