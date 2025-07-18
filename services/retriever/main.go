package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"urlshortener"
	"urlshortener/services"
)

func main() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	grpcServerWrapper := services.NewGrpcWrapper(urlshortener.AppCfg.RetrieveServerAddr)
	RegisterSvc(grpcServerWrapper.Server())

	log.Println("Starting Retriever server on: ", urlshortener.AppCfg.RetrieveServerAddr)

	go func() {
		if err := grpcServerWrapper.Run(); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	<-quit

	log.Println("Data saved. Exiting.")
}
