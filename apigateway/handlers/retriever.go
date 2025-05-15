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
	shortUrl := c.Param("shortUrl")
	req := pb.RetrieveUrlRq{ShortUrl: shortUrl}

	grpcConn := NewGRPCClient(urlshortener.AppCfg.RetrieveServerAddr)
	defer grpcConn.Close()

	client := pb.NewRetrieveServiceClient(grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.RetrieveUrl(ctx, &req)
	if err != nil {
		log.Fatalf("error calling urlshorten service: %v", err)
	}

	log.Printf("Response from Retrieve server : %s", response.GetLongUrl())
	c.JSON(http.StatusCreated, gin.H{"LongUrl:": response.GetLongUrl()})
}
