package db

import (
	"context"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName = os.Getenv("DBNAME")
var MongoURI = os.Getenv("MONGODB_URI")
var Collection = os.Getenv("COLLECTION")

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
