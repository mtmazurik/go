package main

import (
	"context"
    "fmt"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {

}
