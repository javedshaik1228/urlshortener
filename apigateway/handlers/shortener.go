package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"urlshortener"
	pb "urlshortener/proto/genproto/shortenpb"
)

func CallShortenerSvc(c *gin.Context) {
	var req pb.ShortenUrlRq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcConn := NewGRPCConn(urlshortener.AppCfg.ShortenServerAddr)
	defer grpcConn.Close()

	client := pb.NewShortenServiceClient(grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.ShortenUrl(ctx, &req)
	if err != nil {
		log.Fatalf("error calling urlshorten service: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
	}

	log.Printf("Response from Shorten server : %s", response.GetShortUrl())
	c.JSON(http.StatusCreated, gin.H{"ShortUrl:": response.GetShortUrl()})
}
