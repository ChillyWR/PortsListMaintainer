package internal

import (
	"github.com/ChillyWR/PortsListMaintainer/config"
	"github.com/ChillyWR/PortsListMaintainer/proto"
	"google.golang.org/grpc"
)

type ConnectionWithDomainService struct {
	connection    *grpc.ClientConn
	ClientService proto.AddServiceClient
}

func ConnectToDomainService() (*ConnectionWithDomainService, error) {
	var conn ConnectionWithDomainService
	// replace with grpc.WithTransportCredentials(insecure.NewCredentials())
	connection, err := grpc.Dial("localhost:"+config.DomainPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	conn.connection = connection
	conn.ClientService = proto.NewAddServiceClient(conn.connection)

	return &conn, nil
}

func (conn *ConnectionWithDomainService) Close() error {
	err := conn.connection.Close()
	return err
}
