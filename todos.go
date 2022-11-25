package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func insert(db *mongo.Client) {
	todos := db.Database("todos").Collection("todos")
	todo := bson.D{{"description", fmt.Sprintf("todo %s", time.Now().Format("02-01-2006 15:04:05"))}}
	result, err := todos.InsertOne(context.TODO(), todo)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}
