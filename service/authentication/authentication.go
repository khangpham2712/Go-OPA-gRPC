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
		log.Fatalln("Token generation error: " + err.Error())
	}

	var successStatus int32 = 1
	var output *proto.AutOutput = &proto.AutOutput{SuccessStatus: successStatus, Token: accessToken}
	return output, nil
}
