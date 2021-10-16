package api

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) HelloWorld(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("receive message body from client : %s", in.Greeting) //we declared greetring wit small g in .proot file but in generated we have capital G. thus used capital G.
	return &PingMessage{Greeting: "Hello from server side"}, nil
}
