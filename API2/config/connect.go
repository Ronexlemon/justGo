package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Value string = "mongodb+srv://huge65702:huge65702@cluster0.3s02tjq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

func newMongoClient() *mongo.Client {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(Value))
	if err != nil {
		log.Fatal("error while connecting db", err)
	}
	log.Println("Connection successful")
	err = mongoClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal("error while pinging db", err)
	}
	return mongoClient

}
