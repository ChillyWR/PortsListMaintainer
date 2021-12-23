package main

import (
	"github.com/ChillyWR/PortsListMaintainer/ClientAPIService/internal"
)

func main() {
	conn := internal.ConnectToDomainService()
	defer conn.Close()
	internal.StartClientAPI(conn)

}
