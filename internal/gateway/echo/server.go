package echo

import (
	"context"

	"google.golang.org/grpc"
	"log"
	"net"
)

func CreateServer() {
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Prepare in port 9090")

	rpcServer := grpc.NewServer()
	client := &echoServer{}
	RegisterEchoServiceServer(rpcServer, client)
	if rpcServer == nil {
		log.Fatalf("failed to register server: %v", err)
	}
	if err := rpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("Grpc server listening port 9090")
}

type echoServer struct {
}

func (e *echoServer) Echo(context.Context, *StringMessage) (*StringMessage, error) {
	return &StringMessage{Value: "Well done!"}, nil
}
