package main

import (
	"github.com/ChillyWR/PortsListMaintainer/ClientAPIService/internal"
	"log"
)

func main() {
	conn, err := internal.ConnectToDomainService()
	if err != nil {
		log.Fatalf("Failed to connect to domain service: %v", err)
	}
	defer func(conn *internal.ConnectionWithDomainService) {
		err := conn.Close()
		if err != nil {
			log.Println("Failed to close connection with domain service")
		}
	}(conn)
	internal.StartClientAPI(conn)

}
