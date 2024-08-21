package database

import (
	"context"
	"fmt"
	"log"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDB() error {
	// Replace with your MongoDB cloud connection string
	uri := "mongodb+srv://jaddibeh:Jaddibeh2@cluster0.siq35mq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	opts := options.Client().ApplyURI(uri)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Verify the connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	MongoClient = client
	fmt.Println("Successfully connected to MongoDB!")
	return nil
}

func DisconnectDB() {
	if MongoClient != nil {
		err := MongoClient.Disconnect(context.TODO())
		if err != nil {
			log.Printf("Failed to disconnect MongoDB client: %v\n", err)
		} else {
			fmt.Println("Disconnected from MongoDB")
		}
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	if MongoClient == nil {
		panic("MongoClient is nil, make sure ConnectDB is called first")
	}
	return MongoClient.Database("todo").Collection(collectionName)
}
