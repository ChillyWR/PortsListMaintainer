package main

import (
	"github.com/ChillyWR/PortsListMaintainer/DomainService/internal"
	"github.com/ChillyWR/PortsListMaintainer/DomainService/internal/db"
	"log"
)

func main() {
	repo, err := db.NewRepo()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer func(repo *db.Repo) {
		err := repo.CloseConnectionWithDB()
		if err != nil {
			log.Printf("Failed to close connection with DB: %v", err)
		}
	}(repo)
	internal.StartServer(repo)
}
