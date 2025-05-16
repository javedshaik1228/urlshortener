package main

import (
	"context"
	"log"

	pb "urlshortener/proto/genproto/retrievepb"
	"urlshortener/services/retriever/usecase"
	"urlshortener/services/store"

	"google.golang.org/grpc"
)

type RetrieveServer struct {
	pb.UnimplementedRetrieveServiceServer
}

func (s *RetrieveServer) RetrieveUrl(ctx context.Context, req *pb.RetrieveUrlRq) (*pb.RetrieveUrlRs, error) {

	// get nosql client instance
	noSqlClient, err := store.GetNoSQLClient()
	if err != nil {
		log.Printf("Failed to get NoSQLClient: %v", err)
		return nil, err
	}

	longUrl := usecase.UcRetriever(noSqlClient, req.GetShortUrl())
	if longUrl == "" {
		log.Printf("No long URL found for short URL: %s", req.GetShortUrl())
		return nil, nil
	}
	return &pb.RetrieveUrlRs{LongUrl: longUrl}, nil
}

func RegisterSvc(grpcSrv *grpc.Server) {
	pb.RegisterRetrieveServiceServer(grpcSrv, &RetrieveServer{})
}
