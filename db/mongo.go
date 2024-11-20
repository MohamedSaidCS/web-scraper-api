package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitMongoDB() {
	ctx := context.TODO()
	var err error
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+host+":"+port))
	if err != nil {
		log.Panic(err)
	}

	var result bson.M
	if err := Client.Database("go").RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		log.Panic(err)
	}
}
