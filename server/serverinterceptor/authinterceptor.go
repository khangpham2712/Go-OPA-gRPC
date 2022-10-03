package serverinterceptor

import (
	"bytes"
	"context"
	"dummy/opa/opaserver"
	"encoding/json"
	"errors"
	"fmt"
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

	inputRaw := "{\"token\": \"" + accessToken + "\", \"service\": \"" + info.FullMethod + "\"}"

	fmt.Println(inputRaw)

	var input interface{}
	err := json.NewDecoder(bytes.NewBufferString(inputRaw)).Decode(&input)
	if err != nil {
		return nil, err
	}

	fmt.Println(input)

	isAllowed := opaserver.QueryOPAServer(input)
	if !isAllowed {
		return nil, errors.New("Unauthorized")
	}

	return handler(ctx, req)
}

func StreamAuthServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("--> Stream Interceptor:", info.FullMethod)

	return handler(srv, ss)
}
