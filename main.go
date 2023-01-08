package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
const uri = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

func main() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	// Create a new database named "users_details"
	database := client.Database("users_details")

	// Create a new collection named "users" in the "users_details" database
	collection := database.Collection("users")

	// Insert a document into the "users" collection
	doc := map[string]interface{}{
		"Name":    "Hemel Bhai",
		"Address": "Dhaka",
		"Company": "AHOM Limited",
	}
	insertResult, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
