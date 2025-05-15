package store

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const kShortUrlDocKey = "shortUrl"
const kLongUrlDocKey = "longUrl"

type UrlDocument struct {
	ShortUrl  string `bson:"shortUrl"`
	LongUrl   string `bson:"longUrl"`
	CreatedAt string `bson:"createdAt"`
}

func findDocument(keyName string, value string) *UrlDocument {

	NoSQLClient := NewNoSQLClient()
	filter := bson.D{{Key: keyName, Value: value}}

	var result UrlDocument
	err := NoSQLClient.UrlCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err.Error(), filter)
		return nil
	}
	fmt.Printf("Document Found:\n%+v\n", result)
	return &result
}

func pushDocument(newDoc *UrlDocument) {
	NoSQLClient := NewNoSQLClient()
	NoSQLClient.UrlCollection.InsertOne(context.TODO(), newDoc)
}

func InsertDocument(shortUrl string, longUrl string) {
	newDoc := UrlDocument{ShortUrl: shortUrl, LongUrl: longUrl, CreatedAt: fmt.Sprint(time.Now().UnixNano())}
	pushDocument(&newDoc)
}

func FetchDocFromShortUrl(shortUrl string) *UrlDocument {
	return findDocument(kShortUrlDocKey, shortUrl)
}

func FetchDocFromLongUrl(longUrl string) *UrlDocument {
	return findDocument(kLongUrlDocKey, longUrl)
}
