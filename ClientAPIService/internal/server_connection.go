package internal

import (
	"github.com/ChillyWR/PortsListMaintainer/proto"
	"google.golang.org/grpc"
	"log"
)

type ConnectionWithDomainService struct {
	connection    *grpc.ClientConn
	ClientService proto.AddServiceClient
}

func ConnectToDomainService() *ConnectionWithDomainService {
	var conn ConnectionWithDomainService
	connection, err := grpc.Dial("localhost:"+"9999", grpc.WithInsecure()) //TODO: make external config file
	conn.connection = connection
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	conn.ClientService = proto.NewAddServiceClient(conn.connection)

	return &conn
}

func (conn *ConnectionWithDomainService) Close() {
	err := conn.connection.Close()
	if err != nil {
		log.Fatalf("Failed to close connection: %v", err)
	}
}
