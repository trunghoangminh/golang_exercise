package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const URI = "mongodb://localhost:27017"

var ctx = context.TODO()
var databaseName = "school"

func ConnectMongoDB() *mongo.Database {
	// Replace the uri string with your MongoDB deployment's connection string.
	clientOptions := options.Client().ApplyURI(URI)
	var client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	database := client.Database(databaseName)
	fmt.Println("Successfully connected and pinged.")
	return database
}
