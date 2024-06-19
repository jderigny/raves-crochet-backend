package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func OpenConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	context, contextCancelFunc := context.WithCancel(context.Background())

	client, err := mongo.Connect(context, options.Client().ApplyURI("mongodb://root:password@mongodb")) //TODO : move URL & credentials
	if err != nil {
		panic(fmt.Sprintf("MongoDB connection issue %s", err))
	}

	if err = client.Ping(context, readpref.Primary()); err != nil {
		panic(fmt.Sprintf("MongoDB ping issue %s", err))
	}

	fmt.Println("MongoDB connection is opened")

	return client, context, contextCancelFunc
}

func CloseConnection(client *mongo.Client, context context.Context, contextCancelFunc context.CancelFunc) {
	contextCancelFunc()

	if err := client.Disconnect(context); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB connection is closed")
}
