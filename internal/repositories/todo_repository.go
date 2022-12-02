package repositories

import (
	"bot-mongo/internal/configs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TodoRepository interface {
	Insert()
}

type MongoRepository struct {
	db *mongo.Client
}

func (m *MongoRepository) Insert() {
	todos := m.db.Database("todos").Collection("todos")
	todo := bson.D{{"description", fmt.Sprintf("todo %s", time.Now().Format("02-01-2006 15:04:05"))}}
	result, err := todos.InsertOne(context.TODO(), todo)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
}

func NewTodoRepository() TodoRepository {
	mongoDb := configs.OpenConnection()
	defer configs.CloseClientDB()

	repository := &MongoRepository{
		db: mongoDb,
	}

	return repository
}
