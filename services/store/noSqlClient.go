package store

import (
	"context"
	"log"
	"sync"
	"time"

	"urlshortener"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NoSQLClient struct {
	Client        *mongo.Client
	UrlCollection *mongo.Collection
}

var (
	clientInstance *NoSQLClient
	once           sync.Once
)

func GetNoSQLClient() (*NoSQLClient, error) {
	var err error
	once.Do(func() {
		clientInstance, err = newNoSQLClient()
		if err != nil {
			log.Printf("Failed to initialize NoSQLClient: %v", err)
		}
	})
	return clientInstance, err
}

func newNoSQLClient() (*NoSQLClient, error) {
	log.Println("Initializing NoSQLClient...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := initializeMongoClient(ctx, urlshortener.DbCnxUri)
	if err != nil {
		return nil, err
	}

	collection := client.Database(urlshortener.DbCfg.Database).Collection(urlshortener.DbCfg.Collection)

	return &NoSQLClient{
		Client:        client,
		UrlCollection: collection,
	}, nil
}

func initializeMongoClient(ctx context.Context, uri string) (*mongo.Client, error) {
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}

	// Verify the connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Printf("Failed to ping MongoDB: %v", err)
		return nil, err
	}

	return client, nil
}
