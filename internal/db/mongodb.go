package db

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
	"os"
	"sync"
	"time"
)

var mongoDBClient *mongo.Client
var databaseInstance *mongo.Database
var once sync.Once
var defaultCollection *mongo.Collection

func ConnectDB() (*mongo.Client, *mongo.Database, *mongo.Collection) {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		dbUrl := os.Getenv("DATABASE_URL")
		if dbUrl == "" {
			log.Fatalf("No database url found ")
		}

		client, err := mongo.Connect(options.Client().ApplyURI(dbUrl))

		if err != nil {
			log.Fatalf("MongoDB connection failed: %v", err)
		}

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			log.Fatalf("MongoDB ping failed: %v", err)
		}

		log.Println("✅ Connected to MongoDB")

		databaseName := os.Getenv("DATABASE_NAME")

		if databaseName == "" {
			log.Fatalf("No database name found ")
		}

		dbInstance := client.Database(databaseName)

		mongoDBClient = client
		databaseInstance = dbInstance

		collectionName := os.Getenv("COLLECTION_NAME")
		defaultCollection = dbInstance.Collection(collectionName)

	})

	return mongoDBClient, databaseInstance, defaultCollection

}
