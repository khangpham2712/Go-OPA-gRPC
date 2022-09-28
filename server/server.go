package server

import (
	"dummy/proto"

	"dummy/server/serverinterceptor"
	"dummy/service/authentication"
	"dummy/service/multiplication"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunServer() {
	listener, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatalln("Something went wrong: " + err.Error())
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(serverinterceptor.UnaryAuthServerInterceptor),
		grpc.StreamInterceptor(serverinterceptor.StreamAuthServerInterceptor))

	proto.RegisterMultiplicationServer(server, &multiplication.MultiplicationServer{})
	proto.RegisterAuthenticationServer(server, &authentication.AuthenticationServer{})

	err = server.Serve(listener)
	if err != nil {
		log.Fatalln("Something went wrong: " + err.Error())
	}
}
