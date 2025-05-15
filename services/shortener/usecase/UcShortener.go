package usecase

import (
	"fmt"
	"time"
	"urlshortener/services/store"
)

func UcShortener(longUrl string, userid string) string {

	urlData := store.UrlData{LongUrl: longUrl, CreatedAt: fmt.Sprint(time.Now().UnixNano())}

	newShortUrl := RandGenerator(longUrl)
	dataMap := store.GetDataStoreInstance().DataMap

	dataMap[newShortUrl] = urlData

	return newShortUrl
}
