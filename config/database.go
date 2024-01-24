package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	LoadConfig()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(AppConfig.MONGOURI))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
	return client
}

var DB *mongo.Client = ConnectToDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("aurora").Collection(collectionName)
	return collection
}
