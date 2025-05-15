package usecase

import (
	"fmt"
	"urlshortener/services/store"
)

func UcShortener(longUrl string, userid string) string {

	// Check for existing entry
	doc := store.FetchDocFromLongUrl(longUrl)
	if doc != nil {
		fmt.Println("Entry already exists")
		return doc.ShortUrl
	}

	// doc is nil => it is a new entry
	newShortUrl := RandGenerator(longUrl)
	store.InsertDocument(newShortUrl, longUrl)
	return newShortUrl
}
