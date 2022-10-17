package serverinterceptor

import (
	"context"
	opaserver "dummy/opa"
	"sync"

	"time"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var mutex = sync.RWMutex{}

var noAuthorization map[string]bool = map[string]bool{
	"/proto.Authentication/Authenticate": true,
}

var receivedTokens map[string]bool = map[string]bool{}

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

	if isAuth, has := readFromReceivedTokens(accessToken + ".." + info.FullMethod); has {
		if isAuth {
			return handler(ctx, req)
		}

		log.Println("Unauthorized")
		return nil, nil
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

	writeToReceivedTokens(accessToken+".."+info.FullMethod, isAllowed)

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

	if isAuth, has := readFromReceivedTokens(accessToken + ".." + info.FullMethod); has {
		if isAuth {
			return handler(srv, ss)
		}

		log.Println("Unauthorized")
		return nil
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

	writeToReceivedTokens(accessToken+".."+info.FullMethod, isAllowed)

	if !isAllowed {
		log.Println("Unauthorized")
		return nil
	}

	return handler(srv, ss)
}

func readFromReceivedTokens(s string) (bool, bool) {
	mutex.RLock()
	isAllowed, has := receivedTokens[s]
	mutex.RUnlock()
	return isAllowed, has
}

func writeToReceivedTokens(s string, b bool) {
	mutex.Lock()
	receivedTokens[s] = b
	mutex.Unlock()
}
