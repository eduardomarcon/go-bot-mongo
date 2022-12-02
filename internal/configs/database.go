package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var client *mongo.Client

func OpenConnection() *mongo.Client {
	if client != nil {
		return client
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(GetDB().URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseClientDB() {
	if client == nil {
		return
	}
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
