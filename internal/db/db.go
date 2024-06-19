package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func OpenConnection() *mongo.Client {
	context := context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://root:password@mongodb") //TODO : move URL & credentials

	client, err := mongo.Connect(context, clientOptions)
	if err != nil {
		panic(fmt.Sprintf("MongoDB connection issue %s", err))
	}

	if err = client.Ping(context, nil); err != nil {
		panic(fmt.Sprintf("MongoDB ping issue %s", err))
	}

	fmt.Println("MongoDB connection is opened")

	return client
}
