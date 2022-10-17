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

var receivedTokens map[string]int64 = map[string]int64{}

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

	if e, has := readFromReceivedTokens(accessToken + ".." + info.FullMethod); has {
		if time.Now().Unix() <= e {
			return handler(ctx, req)
		}

		delete(receivedTokens, accessToken+".."+info.FullMethod)
		log.Println("Token expired")
		return nil, nil
	}

	input := map[string]string{
		"token":   accessToken,
		"service": info.FullMethod,
	}
	//this code block is used to log the OPA execution time
	start := time.Now()
	isAllowed, exp := opaserver.QueryOPAServer(input)
	duration := time.Since(start)
	log.Println("OPA query duration:", duration)
	//

	writeToReceivedTokens(accessToken+".."+info.FullMethod, exp)

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

	if e, has := readFromReceivedTokens(accessToken + ".." + info.FullMethod); has {
		if time.Now().Unix() <= e {
			return handler(srv, ss)
		}

		delete(receivedTokens, accessToken+".."+info.FullMethod)
		log.Println("Token expired")
		return nil
	}

	input := map[string]string{
		"token":   accessToken,
		"service": info.FullMethod,
	}
	//this code block is used to log the OPA execution time
	start := time.Now()
	isAllowed, exp := opaserver.QueryOPAServer(input)
	duration := time.Since(start)
	log.Println("OPA query duration:", duration)
	//

	writeToReceivedTokens(accessToken+".."+info.FullMethod, exp)

	if !isAllowed {
		log.Println("Unauthorized")
		return nil
	}

	return handler(srv, ss)
}

func readFromReceivedTokens(s string) (int64, bool) {
	mutex.RLock()
	exp, has := receivedTokens[s]
	mutex.RUnlock()
	return exp, has
}

func writeToReceivedTokens(s string, exp int64) {
	mutex.Lock()
	receivedTokens[s] = exp
	mutex.Unlock()
}
