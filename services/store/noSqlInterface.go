package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	kLongUrlDocKey  = "longUrl"
	kShortUrlDocKey = "shortUrl"
)

type UrlDocument struct {
	ShortUrl  string `bson:"shortUrl"`
	LongUrl   string `bson:"longUrl"`
	CreatedAt string `bson:"createdAt"`
}

func findDocument(ctx context.Context, noSqlClient *NoSQLClient, keyName, value string) (*UrlDocument, error) {
	filter := bson.D{{Key: keyName, Value: value}}

	var result UrlDocument
	err := noSqlClient.UrlCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No document found for filter: %v", filter)
			return nil, nil
		}
		log.Printf("Error finding document: %v, filter: %v", err, filter)
		return nil, err
	}

	log.Printf("Document found: %+v", result)
	return &result, nil
}

func pushDocument(ctx context.Context, noSqlClient *NoSQLClient, newDoc *UrlDocument) error {
	_, err := noSqlClient.UrlCollection.InsertOne(ctx, newDoc)
	if err != nil {
		log.Printf("Error inserting document: %v, document: %+v", err, newDoc)
		return err
	}
	log.Printf("Document inserted successfully: %+v", newDoc)
	return nil
}

func (noSqlClient *NoSQLClient) InsertDocument(shortUrl, longUrl string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newDoc := UrlDocument{
		ShortUrl:  shortUrl,
		LongUrl:   longUrl,
		CreatedAt: fmt.Sprint(time.Now().UnixNano()),
	}
	return pushDocument(ctx, noSqlClient, &newDoc)
}

func (noSqlClient *NoSQLClient) FetchDocFromShortUrl(shortUrl string) (*UrlDocument, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return findDocument(ctx, noSqlClient, kShortUrlDocKey, shortUrl)
}

func (noSqlClient *NoSQLClient) FetchDocFromLongUrl(longUrl string) (*UrlDocument, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return findDocument(ctx, noSqlClient, kLongUrlDocKey, longUrl)
}
