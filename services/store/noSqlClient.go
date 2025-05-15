package store

import (
	"context"
	"time"

	"urlshortener"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NoSQLClient struct {
	Client        *mongo.Client
	UrlCollection *mongo.Collection
}

func NewNoSQLClient() *NoSQLClient {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(urlshortener.DbCnxUri)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(err)
	}

	collection := client.Database(urlshortener.DbCfg.Database).Collection(urlshortener.DbCfg.Collection)

	return &NoSQLClient{
		Client:        client,
		UrlCollection: collection,
	}
}
