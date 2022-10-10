package serverinterceptor

import (
	"context"
	opaserver "dummy/opa"
	"errors"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryAuthServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("--> Unary Interceptor:", info.FullMethod)

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

	isAllowed := opaserver.QueryOPAServer(input)
	if !isAllowed {
		log.Println("Unauthorized")
		// return nil, errors.New("Unauthorized")
		// return proto.Output{ErrorCode: 1, Res: 0}, nil
		return nil, nil
	}

	return handler(ctx, req)
}

func StreamAuthServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("--> Stream Interceptor:", info.FullMethod)

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

	isAllowed := opaserver.QueryOPAServer(input)
	if !isAllowed {
		return errors.New("Unauthorized")
	}

	return handler(srv, ss)
}
