package usecase

import (
	"math/rand"
	"time"
)

const kShortUrlLength = 6
const kCharSet = "ABCDEFGHIJKLMNOPQRSTUVXWYZabcdefghijklmnopqrstuvxwyz0123456789"

func RandGenerator(longUrl string) string {
	shortUrl := make([]byte, kShortUrlLength)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range kShortUrlLength {
		idx := r.Intn(len(kCharSet))
		shortUrl[i] = kCharSet[idx]
	}
	return string(shortUrl)
}
