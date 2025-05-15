package main

import (
	"context"

	pb "urlshortener/proto/genproto/shortenpb"
	uc "urlshortener/services/shortener/usecase"

	"google.golang.org/grpc"
)

type ShortenServer struct {
	pb.UnimplementedShortenServiceServer
}

func (s *ShortenServer) ShortenUrl(ctx context.Context, req *pb.ShortenUrlRq) (*pb.ShortenUrlRs, error) {
	shortUrl := uc.UcShortener(req.GetLongUrl(), req.GetUserId())
	return &pb.ShortenUrlRs{ShortUrl: shortUrl}, nil
}

func RegisterSvc(grpcSrv *grpc.Server) {
	pb.RegisterShortenServiceServer(grpcSrv, &ShortenServer{})
}
