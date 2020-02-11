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
	Name 	string
	Age		int
	City string
}

func main() {
	// Creat new Trainer Structs
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB - .TODO returns a non-nil empty context 
	// becuase the surrounding function has not accepted a context paramter
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection - see if server has been found and connected 
	// using the Ping method
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to MongoDB!")

	collection := client.Database("test").Collection("trainers")
	filter := bson.D{{"name", "Ash"}}

	update := bson.D{
	    {"$inc", bson.D{
	        {"age", 1},
	    }},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
	    log.Fatal(err)
	}

fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)


	trainers := []interface{}{misty, brock, ash}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
	    log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	err = client.Disconnect(context.TODO())

	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}