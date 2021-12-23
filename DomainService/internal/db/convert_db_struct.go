package db

import "github.com/ChillyWR/PortsListMaintainer/proto"

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

func convertToProtoPortStruct(port *Port) *proto.Port {
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

func convertToPortStruct(port *proto.Port) *Port {
	protoPort := Port{
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
