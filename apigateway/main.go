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

	api := router.Group("/urlshorterner")
	{
		api.GET("/:shortUrl", handlers.CallRetrieverSvc)
		api.POST("/shorten", handlers.CallShortenerSvc)
	}

	router.Run(urlshortener.AppCfg.GatewayServerAddr)
}
