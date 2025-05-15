package main

import (
	"urlshortener"

	"urlshortener/apigateway/handlers"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	router := gin.Default()

	router.GET("/urlshorterner/:shortUrl", func(c *gin.Context) {
		handlers.CallRetrieverSvc(c)
	})

	router.POST("/urlshorterner/shorten", func(c *gin.Context) {
		handlers.CallShortenerSvc(c)
	})

	router.Run(urlshortener.AppCfg.GatewayServerAddr)
}
