package gRPC

import (
	"context"
	"github.com/ChillyWR/PortsListMaintainer/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct{}

func (s *server) ListPort(ctx context.Context, request *proto.RequestToListPort) (*proto.Port, error) {
	//TODO: get ports from database
	return &proto.Port{Port: "test port"}, nil
}

func (s *server) AddPort(ctx context.Context, port *proto.Port) (*proto.Success, error) {
	log.Printf("Got port %s", port.Port)
	//TODO: upsert ports to database
	return &proto.Success{}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":"+"9999") //TODO: make external config file
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterAddServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
