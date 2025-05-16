package main

import (
	"context"
	"log"

	pb "urlshortener/proto/genproto/shortenpb"
	uc "urlshortener/services/shortener/usecase"
	"urlshortener/services/store"

	"google.golang.org/grpc"
)

type ShortenServer struct {
	pb.UnimplementedShortenServiceServer
}

func (s *ShortenServer) ShortenUrl(ctx context.Context, req *pb.ShortenUrlRq) (*pb.ShortenUrlRs, error) {
	// get nosql client instance
	noSqlClient, err := store.GetNoSQLClient()
	if err != nil {
		log.Printf("Failed to get NoSQLClient: %v", err)
		return nil, err
	}

	shortUrl := uc.UcShortener(noSqlClient, req.GetLongUrl(), req.GetUserId())
	return &pb.ShortenUrlRs{ShortUrl: shortUrl}, nil
}

func RegisterSvc(grpcSrv *grpc.Server) {
	pb.RegisterShortenServiceServer(grpcSrv, &ShortenServer{})
}
