package server

import (
	opaserver "dummy/opa"
	"dummy/proto"
	"dummy/server/serverinterceptor"

	"dummy/service/authentication"
	"dummy/service/multiplication"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var protocol string = "tcp"
var port string = "50051"

func RunGRPCServer() {
	listener, err := net.Listen(protocol, ":"+port)
	if err != nil {
		log.Fatalln("Listener error: " + err.Error())
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(serverinterceptor.UnaryAuthServerInterceptor),
		grpc.StreamInterceptor(serverinterceptor.StreamAuthServerInterceptor))

	proto.RegisterMultiplicationServer(server, &multiplication.MultiplicationServer{})
	proto.RegisterAuthenticationServer(server, &authentication.AuthenticationServer{})

	reflection.Register(server)

	opaserver.RegisterOPAQuery()

	err = server.Serve(listener)
	if err != nil {
		log.Fatalln("Server error: " + err.Error())
	} else {
		log.Println("gRPC server is listening on port " + port)
	}
}
