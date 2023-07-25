package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a global variable to hold the MongoDB collection
var Con *mongo.Collection

// Initialize the MongoDB connection and collection in your main function or init function
func init() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://mritun:8011501382@mritundbcluster.tte3tkm.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	db := client.Database("andybrandproject")
	Con = db.Collection("users")
}
