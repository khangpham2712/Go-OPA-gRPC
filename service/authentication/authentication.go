package authentication

import (
	"context"
	"dummy/proto"
	"dummy/token"
	"log"
)

type AuthenticationServer struct {
	proto.UnimplementedAuthenticationServer
}

func (server *AuthenticationServer) Authenticate(c context.Context, input *proto.AutInput) (*proto.AutOutput, error) {
	username := input.GetUsername()
	role := input.GetRole()
	accessToken, err := token.Generate(username, role, "dummy")
	if err != nil {
		log.Fatalln(err.Error())
	}

	var output *proto.AutOutput = &proto.AutOutput{Token: accessToken}
	return output, nil
}
