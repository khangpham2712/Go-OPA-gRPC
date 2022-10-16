package serverinterceptor

import (
	"context"
	opaserver "dummy/opa"

	"time"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var noAuthorization map[string]bool = map[string]bool{
	"/proto.Authentication/Authenticate": true,
}

func UnaryAuthServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("--> Unary Interceptor:", info.FullMethod)

	if _, ok := noAuthorization[info.FullMethod]; ok {
		return handler(ctx, req)
	}

	var accessToken string = ""
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if accT, has := md["access_token"]; has {
			accessToken = accT[0]
		}
	}

	input := map[string]string{
		"token":   accessToken,
		"service": info.FullMethod,
	}
	//this code block is used to log the OPA execution time
	start := time.Now()
	isAllowed := opaserver.QueryOPAServer(input)
	duration := time.Since(start)
	log.Println("OPA query duration:", duration)
	//
	if !isAllowed {
		log.Println("Unauthorized")
		return nil, nil
	}

	return handler(ctx, req)
}

func StreamAuthServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("--> Stream Interceptor:", info.FullMethod)

	if _, ok := noAuthorization[info.FullMethod]; ok {
		return handler(srv, ss)
	}

	var accessToken string = ""
	if md, ok := metadata.FromIncomingContext(ss.Context()); ok {
		if accT, has := md["access_token"]; has {
			accessToken = accT[0]
		}
	}

	input := map[string]string{
		"token":   accessToken,
		"service": info.FullMethod,
	}
	//this code block is used to log the OPA execution time
	start := time.Now()
	isAllowed := opaserver.QueryOPAServer(input)
	duration := time.Since(start)
	log.Println("OPA query duration:", duration)
	//
	if !isAllowed {
		log.Println("Unauthorized")
		return nil
	}

	return handler(srv, ss)
}
