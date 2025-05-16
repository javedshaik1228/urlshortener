package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"urlshortener"
	pb "urlshortener/proto/genproto/retrievepb"
)

func CallRetrieverSvc(c *gin.Context) {
	shortUrlParam := c.Param("shortUrl")
	req := pb.RetrieveUrlRq{ShortUrl: shortUrlParam}

	grpcConn, err := NewGRPCConn(urlshortener.AppCfg.RetrieveServerAddr)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Failed to connect to Retrieve service"})
		return
	}

	defer grpcConn.Close()

	client := pb.NewRetrieveServiceClient(grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.RetrieveUrl(ctx, &req)
	if err != nil {
		log.Fatalf("error calling urlshorten service: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
	}

	log.Printf("Response from Retrieve server : %s", response.GetLongUrl())
	c.JSON(http.StatusCreated, gin.H{"LongUrl:": response.GetLongUrl()})
}
