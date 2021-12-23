package main

import (
	"github.com/ChillyWR/PortsListMaintainer/DomainService/internal/db"
	"github.com/ChillyWR/PortsListMaintainer/DomainService/internal/gRPC"
)

func main() {
	repo := db.ConnectToDBWithGorm()
	defer repo.CloseConnectionWithDB()
	gRPC.StartServer(repo)
}
