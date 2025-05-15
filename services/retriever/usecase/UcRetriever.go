package usecase

import (
	"urlshortener/services/store"
)

func UcRetriever(shortUrl string) string {
	longUrl := store.FetchDocFromShortUrl(shortUrl).LongUrl
	return longUrl
}
