package main

import (
	"context"

	pb "urlshortener/proto/genproto/retrievepb"
	uc "urlshortener/services/retriever/usecase"

	"google.golang.org/grpc"
)

type RetrieveServer struct {
	pb.UnimplementedRetrieveServiceServer
}

func (s *RetrieveServer) RetrieveUrl(ctx context.Context, req *pb.RetrieveUrlRq) (*pb.RetrieveUrlRs, error) {
	longUrl := uc.UcRetriever(req.GetShortUrl())
	return &pb.RetrieveUrlRs{LongUrl: longUrl}, nil
}

func RegisterSvc(grpcSrv *grpc.Server) {
	pb.RegisterRetrieveServiceServer(grpcSrv, &RetrieveServer{})
}
