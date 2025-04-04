package Db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var DB *mongo.Database

func OpenConnection() {
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Cannot connect to mongo " + err.Error())
	}

	// Ping the primary node to verify connection is established.
	ctxPing, cancelPing := context.WithTimeout(context.Background(), 5*time.Second) // Shorter timeout for ping

	defer cancelPing()

	err = client.Ping(ctxPing, readpref.Primary())

	if err != nil {
		// Disconnect if ping fails
		_ = client.Disconnect(context.Background()) // Attempt cleanup
		fmt.Errorf("cannot connect to mongo (ping failed): %w", err)
		return
	}

	DB = client.Database("todo")

}
