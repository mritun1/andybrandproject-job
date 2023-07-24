package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg MongoInstance

// ---------------------------------------------------------------------
//
//	COLLECTION NAME AND MONGODB URI - START
//
// ---------------------------------------------------------------------
const dbName = "andybrandproject"
const MongoURI = "mongodb+srv://mritun:8011501382@mritundbcluster.tte3tkm.mongodb.net/" + dbName

// ---------------------------------------------------------------------
//
//	COLLECTION NAME AND MONGODB URI - END
//
// ---------------------------------------------------------------------
// ---------------------------------------------------------------------
//
//	START CONNECTION TO MONGODB - START
//
// ---------------------------------------------------------------------
func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	db := client.Database(dbName)
	if err != nil {
		return err
	}
	Mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

// ---------------------------------------------------------------------
//
//	START CONNECTION TO MONGODB - END
//
// ---------------------------------------------------------------------
