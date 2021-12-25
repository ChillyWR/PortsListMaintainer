package internal

import (
	"context"
	"github.com/ChillyWR/PortsListMaintainer/DomainService/internal/db"
	"github.com/ChillyWR/PortsListMaintainer/config"
	"github.com/ChillyWR/PortsListMaintainer/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	repo *db.Repo
}

func StartServer(dbConnection *db.Repo) {
	srv := server{repo: dbConnection}
	lis, err := net.Listen("tcp", ":"+config.DomainPort) //TODO: make external config file
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterAddServiceServer(grpcServer, &srv)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) ListPort(_ context.Context, request *proto.RequestToListPorts) (*proto.Ports, error) {
	log.Printf("Got request to list ports")
	id := request.Filters.GetId()
	idLower := request.Filters.GetIdLower()
	idUpper := request.Filters.GetIdUpper()
	filterString := request.Filters.GetFilterString()
	result, err := s.repo.ListPorts(id, idLower, idUpper, &filterString)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *server) UpsertPorts(_ context.Context, ports *proto.Ports) (*proto.DefaultResponse, error) {
	portsList := ports.GetPorts()
	savedIDs := make([]int64, len(portsList))
	for i, port := range portsList {
		savedIDs[i] = port.GetID()
		log.Printf("Got ports %s", port.GetName())
	}
	s.repo.UpsertPorts(ports)
	return &proto.DefaultResponse{Success: true, SavedIdRecord: savedIDs}, nil
}
