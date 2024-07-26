package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBSetup()

func DBSetup() *mongo.Client {
	// Create a new mongodb client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// Connect to client mongodb
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping client using an empty context to see if the connection failed
	client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb")
		return nil
	}

	fmt.Println("Successfully connected to mongodb")
	return client
}

func CollectionData(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := &mongo.Collection{}
	collection = client.Database("Ecommerce").Collection(collectionName)

	return collection
}
