package main

import (
	"fmt"
	"log"
	"net"
	"personal_projects/Go-GRPC/api"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello Server.main.go")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := api.server{}
	grpcServer := grpc.NewServer()

	api.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
