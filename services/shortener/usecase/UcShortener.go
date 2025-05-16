package usecase

import (
	"fmt"
	"log"
	"urlshortener/services/store"
)

func UcShortener(noSqlClient *store.NoSQLClient, longUrl string, userid string) string {

	// Check for existing entry
	doc, err := noSqlClient.FetchDocFromLongUrl(longUrl)
	if err != nil {
		log.Printf("Error fetching document for longUrl %s: %v", longUrl, err)
		return ""
	}

	if doc != nil {
		fmt.Println("Entry already exists")
		return doc.ShortUrl
	}

	// doc is nil => it is a new entry
	newShortUrl := RandGenerator(longUrl)
	noSqlClient.InsertDocument(newShortUrl, longUrl)
	return newShortUrl
}
