package mongodb

import (
	"context"
	"fmt"
	"golang_api/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBConnector() *mongo.Client {
	cfg := config.LoadConfig()
	clientOptions := options.Client().ApplyURI(cfg.GetString("MONGODB_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
