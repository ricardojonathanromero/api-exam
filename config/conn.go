package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// CreateSession create a session in Mongo DB
func CreateSession(host string) (*mongo.Client, error) {

	var session *mongo.Client

	// connect to mongo db
	uri := options.Client().ApplyURI(host)
	session, err := mongo.Connect(context.TODO(), uri)

	if err != nil {
		fmt.Println("Error into CreateSessionFunction - Init connection\n", err)
		return session, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = session.Connect(ctx)

	// ping to host to validate the connection
	err = session.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		fmt.Println("Error into CreateSessionFunction - Ping\n", err)
		return session, err
	}

	return session, err
}

// CloseClient close mongo client
func CloseClient(client *mongo.Client) {
	if client != nil {
		err := client.Disconnect(context.Background())
		if err != nil {
			fmt.Println("Can not close client: ", err)
		}
	}
}
