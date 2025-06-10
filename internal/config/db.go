package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	UserCollection *mongo.Collection
)

func InitDB() {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	println()

	db := client.Database(os.Getenv("MONGODB_NAME"))
	UserCollection = db.Collection("users")

}
