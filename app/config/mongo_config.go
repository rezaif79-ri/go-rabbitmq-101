package config

import (
	"context"
	"time"

	envutil "gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/env_util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenMongoBookDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get mongo connection from ENV
	mongoHost := envutil.GetEnv("MONGO_HOST", "localhost")
	mongoPort := envutil.GetEnv("MONGO_PORT", "27027")
	mongoProtocol := envutil.GetEnv("MONGO_PROTOCOL", "mongodb")
	mongoCreds := envutil.GetEnv("MONGO_CREDS", "")

	connUri := mongoProtocol + "://" + mongoCreds + mongoHost + ":" + mongoPort
	clientOptions := options.Client()
	clientOptions.ApplyURI(connUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return client.Database("books_database"), nil
}
