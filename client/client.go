package main

import (
	"fmt"
	"log"
	"personal_projects/Go-GRPC/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	response, err := c.HelloWorld(context.Background(), &api.PingMessage{Greeting: "hello from client side!!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	fmt.Println("response from server revceived is %s", response.Greeting)

}
