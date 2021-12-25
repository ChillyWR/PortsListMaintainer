package pkg

import (
	"github.com/ChillyWR/PortsListMaintainer/config"
	"github.com/ChillyWR/PortsListMaintainer/proto"
)

//int64 ID = 1;
//string PortName = 2;
//bool IsActive = 3;
//string Company = 4;
//string Email = 5;
//string Phone = 6;
//string Address = 7;
//string About = 8;
//string Registered = 9;
//float Latitude = 10;
//float Longitude = 11;

//TODO: remake with interfaces

func ConvertToProtoPortStruct(port *config.Port) *proto.Port {
	protoPort := proto.Port{
		ID:         port.ID,
		Name:       port.Name,
		IsActive:   port.IsActive,
		Company:    port.Company,
		Email:      port.Email,
		Phone:      port.Phone,
		Address:    port.Address,
		About:      port.About,
		Registered: port.Registered,
		Latitude:   port.Latitude,
		Longitude:  port.Longitude,
	}

	return &protoPort
}

func ConvertToPortStruct(protoPort *proto.Port) *config.Port {
	port := config.Port{
		ID:         protoPort.ID,
		Name:       protoPort.Name,
		IsActive:   protoPort.IsActive,
		Company:    protoPort.Company,
		Email:      protoPort.Email,
		Phone:      protoPort.Phone,
		Address:    protoPort.Address,
		About:      protoPort.About,
		Registered: protoPort.Registered,
		Latitude:   protoPort.Latitude,
		Longitude:  protoPort.Longitude,
	}

	return &port
}
