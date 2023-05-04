package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var BookingCollection *mongo.Collection
var UserCollection *mongo.Collection
var Ctx = context.TODO()

func Init(url string) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	BookingCollection = client.Database("maraka").Collection("booking")
	UserCollection = client.Database("maraka").Collection("users")
}
