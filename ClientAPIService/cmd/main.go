package main

import (
	"context"
	"github.com/ChillyWR/PortsListMaintainer/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:"+"9999", grpc.WithInsecure()) //TODO: make external config file
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}(connection)

	client := proto.NewAddServiceClient(connection)

	port, err := client.ListPort(context.Background(), &proto.RequestToListPort{})
	if err != nil {
		log.Fatalf("Failed to call AddPort method:, %v", err)
	}
	log.Printf("Got response ports: %s", port.Port)

	success, err := client.AddPort(context.Background(), &proto.Port{Port: "1111"})
	if err != nil {
		log.Fatalf("Failed to call AddPort method:, %v", err)
	}
	_ = success

	//g := gin.Default()
	//g.GET("/ListPort/:RequestToListPort", func(ctx *gin.Context) {
	//	RequestToListPort := ctx.Param("RequestToListPort")
	//	_ = RequestToListPort //TODO: parse received nil-data?
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
	//		_ = response //TODO: parse received nil-data?
	//		ctx.JSON(http.StatusOK, gin.H{
	//			"result": "Successfully upsert received port",
	//		})
	//	} else {
	//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	}
	//})
	////
	//if err := g.Run(":" + "9999"); err != nil { //TODO: make external config file
	//	log.Fatalf("Failed to run server: %v", err)
	//}

}
