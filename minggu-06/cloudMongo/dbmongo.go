package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Define the mongodb client URL
	var uri = "mongodb://localhost:27017"

	// Establish the connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Create go routine to defer the closure
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("members").Collection("scoreCollection")
	docs := []interface{}{
		bson.D{{"name", "Jake"}, {"score", 9.7}},
		bson.D{{"name", "Jay"}, {"score", 9.5}},
		bson.D{{"name", "Hee"}, {"score", 7.6}},
		bson.D{{"name", "Johnny"}, {"score", 7.5}},
	}

	// insertMany
	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}
	// end insertMany

	// When you run this file, it should print:
	// Document inserted with ID: ObjectID("...")
	for _, id := range result.InsertedIDs {
		fmt.Printf("\t%s\n", id)
	}

}
