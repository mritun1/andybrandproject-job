package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "andybrandproject"
const MongoURI = "mongodb+srv://mritun:8011501382@mritundbcluster.tte3tkm.mongodb.net/"
const Collection = "users"

// Define a global variable to hold the MongoDB collection
var Con *mongo.Collection

// Initialize the MongoDB connection and collection in your main function or init function
func init() {
	clientOptions := options.Client().ApplyURI(MongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	db := client.Database(dbName)
	Con = db.Collection(Collection)
}
