package usecase

import (
	"log"
	"urlshortener/services/store"
)

func UcRetriever(noSqlClient *store.NoSQLClient, shortUrl string) string {
	doc, err := noSqlClient.FetchDocFromShortUrl(shortUrl)
	if err != nil {
		log.Printf("Error fetching document for shortUrl %s: %v", shortUrl, err)
		return ""
	}
	if doc == nil {
		log.Printf("No document found for shortUrl %s", shortUrl)
		return ""
	}
	return doc.LongUrl
}
