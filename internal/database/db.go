package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var dbName string
var bookMarksCollection *mongo.Collection
var userCollection *mongo.Collection

func ConnectDatabase(url, name string) {
	dbName = name
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Ping the database to check if the connection is successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	bookMarksCollection = client.Database(dbName).Collection("bookmarks")
	log.Printf(bookMarksCollection.Name())
	userCollection = client.Database(dbName).Collection("users")
	log.Println("Connected to MongoDB!")
}

func GetUserCollection() *mongo.Collection {
	if userCollection == nil {
		log.Fatalf("Database connection is not initialized")
	}
	return userCollection
}

func GetBookmarksCollection() *mongo.Collection {
	if bookMarksCollection == nil {
		log.Fatalf("Database connection is not initializeddd")
	}
	return bookMarksCollection
}

func DisconnectDatabase() {
	if client != nil {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from database: %v", err)
		}
	}
}
