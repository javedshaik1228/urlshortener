package usecase

import (
	"fmt"
	"urlshortener/services/store"
)

func UcRetriever(shortUrl string) string {

	dataMap := store.GetDataStoreInstance().DataMap

	longUrl, ok := dataMap[shortUrl]
	if !ok {
		fmt.Println("Mapped url not found for: ", shortUrl)
	}

	return longUrl.LongUrl
}
