package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 1. create a client and connect to it
// 2. ping with client, if all good then connection good
func DBSet() *mongo.Client {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	// if err != nil {
	// 	log.Fatalf("Error while setting up mongoClient\n", err)
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// err = client.Connect(ctx)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("error while connecting to mongo client", err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("Failed to connect to mongodb, ping failed")
		return nil
	}

	fmt.Println("Successfully connected to mongodb client")
	return client
}

var Client *mongo.Client = DBSet()  

// bcz we have 2 databases (or tables) so two funcs
func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return productCollection
}
