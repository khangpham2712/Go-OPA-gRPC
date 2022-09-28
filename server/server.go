package server

import (
	"context"
	"dummy/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedMultiplicationServer
}

func (server *Server) Mul(c context.Context, input *proto.Input) (*proto.Output, error) {
	a, b := input.GetA(), input.GetB()
	result := a * b
	var output *proto.Output = &proto.Output{Res: result}
	time.Sleep(10 * time.Second)
	return output, nil
}

func RunServer() {
	listener, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatalln("Something went wrong: " + err.Error())
	}

	server := grpc.NewServer()
	proto.RegisterMultiplicationServer(server, &Server{})

	err = server.Serve(listener)
	if err != nil {
		log.Fatalln("Something went wrong: " + err.Error())
	}

	log.Printf("Server is listening on port 50000")
}
