package serverinterceptor

import (
	"context"
	"dummy/token"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var privileges map[string][]string = map[string][]string{
	"/proto.Multiplication/Multiply": {"admin"},
}

func UnaryAuthServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("--> Unary Interceptor:", info.FullMethod)

	accessibleRoles, ok := privileges[info.FullMethod]
	if !ok {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("Missing metadata")
	}

	accT, ok := md["access_token"]
	if !ok {
		return nil, errors.New("Missing access token")
	}

	accessToken := accT[0]
	userClaims, err := token.Verify(accessToken)
	if err != nil {
		return nil, err
	}

	role := userClaims.Role
	for _, v := range accessibleRoles {
		if role == v {
			return handler(ctx, req)
		}
	}

	return nil, errors.New("Unauthorized")
}

func StreamAuthServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("--> Stream Interceptor:", info.FullMethod)

	accessibleRoles, ok := privileges[info.FullMethod]
	if !ok {
		return handler(srv, ss)
	}

	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return errors.New("Missing metadata")
	}

	accT, ok := md["access_token"]
	if !ok {
		return errors.New("Missing access token")
	}

	accessToken := accT[0]
	userClaims, err := token.Verify(accessToken)
	if err != nil {
		return err
	}

	role := userClaims.Role
	for _, v := range accessibleRoles {
		if role == v {
			return handler(srv, ss)
		}
	}

	return errors.New("Unauthorized")
}
