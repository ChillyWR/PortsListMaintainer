package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ChillyWR/PortsListMaintainer/ClientAPIService/pkg"
	"github.com/ChillyWR/PortsListMaintainer/config"
	"github.com/ChillyWR/PortsListMaintainer/proto"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ClientAPI struct {
	domainConnection *ConnectionWithDomainService
}

func (c *ClientAPI) HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/list_ports", c.ListPorts).Methods("GET")
	myRouter.HandleFunc("/list_ports/{FilterString}", c.ListPorts).Methods("GET")
	myRouter.HandleFunc("/list_port/{Id}", c.ListPorts).Methods("GET")
	myRouter.HandleFunc("/list_ports/{IdLower}/{IdUpper}", c.ListPorts).Methods("GET")
	myRouter.HandleFunc("/upload_ports", c.UploadPorts).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+config.ClientPort, myRouter))
}

func (c *ClientAPI) ListPorts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: ListPorts")
	vars := mux.Vars(r)
	Id, IdLower, IdUpper, FilterString := pkg.ParseFilters(vars)
	ports, err := c.domainConnection.ClientService.ListPort(context.Background(), &proto.RequestToListPorts{Filters: &proto.FilterVars{
		Id:           Id,
		IdLower:      IdLower,
		IdUpper:      IdUpper,
		FilterString: FilterString,
	}})
	if err != nil {
		log.Printf("Failed to call ClientService.ListPort method:, %v", err)
	}
	if len(ports.GetPorts()) == 0 {
		_, err := fmt.Fprintln(w, "None found")
		if err != nil {
			log.Printf("Failed to write response: %v", err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(ports.Ports)
		if err != nil {
			log.Printf("Failed to encode ports: %v\n -> Failed to write responce", err)
		}
	}
}

//func (c *ClientAPI) UploadPort(w http.ResponseWriter, r *http.Request) {
//	log.Println("Endpoint Hit: UploadPort")
//	w.Header().Set("Content-Type", "application/json")
//	var port proto.Port
//	_ = json.NewDecoder(r.Body).Decode(&port)
//	success, err := c.domainConnection.ClientService.UpsertPort(context.Background(), &port)
//	if success.GetSuccess() == false || err != nil {
//		log.Fatalf("Failed to call UpsertPort method: %v", err)
//	} else {
//		_, err := fmt.Fprintf(w, "Successfully saved record with id %v", success.GetSavedIdRecord())
//		if err != nil {
//			log.Fatalf("Failed to write response: %v", err)
//		}
//	}
//}

func (c *ClientAPI) UploadPorts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: UploadPorts")
	w.Header().Set("Content-Type", "application/json")
	var ports []*proto.Port
	_ = json.NewDecoder(r.Body).Decode(&ports)
	success, err := c.domainConnection.ClientService.UpsertPorts(context.Background(), &proto.Ports{Ports: ports})
	if !success.Success || err != nil {
		log.Fatalf("Failed to call UpsertPort method:, %v", err)
	} else {
		_, err := fmt.Fprintf(w, "Successfully saved record with id %v", success.GetSavedIdRecord())
		if err != nil {
			log.Fatalf("Failed to write response: %v", err)
		}
	}
}

func StartClientAPI(conn *ConnectionWithDomainService) {
	client := ClientAPI{domainConnection: conn}
	client.HandleRequests()
}

//g := gin.Default()
//g.GET("/ListPort/:RequestToListPort", func(ctx *gin.Context) {
//	RequestToListPort := ctx.Param("RequestToListPort")
//	_ = RequestToListPort
//	request := &proto.RequestToListPort{}
//	if response, err := client.ListPort(ctx, request); err == nil {
//		ctx.JSON(http.StatusOK, gin.H{
//			"result": fmt.Sprint(response.Port),
//		})
//	} else {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//	}
//})
//g.GET("/AddPort/:Port", func(ctx *gin.Context) {
//	Port := ctx.Param("Port")
//	request := &proto.Port{Port: Port}
//	if response, err := client.AddPort(ctx, request); err == nil {
//		_ = response
//		ctx.JSON(http.StatusOK, gin.H{
//			"result": "Successfully upsert received port",
//		})
//	} else {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//	}
//})
////
//if err := g.Run(":" + "9999"); err != nil {
//	log.Fatalf("Failed to run server: %v", err)
//}
