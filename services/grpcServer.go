package services

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCWrapper struct {
	addr       string
	grpcServer *grpc.Server
}

// Constructor
func NewGrpcWrapper(addr string) *gRPCWrapper {
	return &gRPCWrapper{
		addr:       addr,
		grpcServer: grpc.NewServer(),
	}
}

// Getter
func (srvWrapper *gRPCWrapper) Server() *grpc.Server {
	return srvWrapper.grpcServer
}

func (srvWrapper *gRPCWrapper) Run() error {
	listener, err := net.Listen("tcp", srvWrapper.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return srvWrapper.grpcServer.Serve(listener)
}
