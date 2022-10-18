package multiplication

import (
	"context"
	"dummy/proto"
)

type MultiplicationServer struct {
	proto.UnimplementedMultiplicationServer
}

func (server *MultiplicationServer) Multiply(c context.Context, input *proto.Input) (*proto.Output, error) {
	a, b := input.GetA(), input.GetB()
	result := a * b
	var successStatus int32 = 1
	var output *proto.Output = &proto.Output{SuccessStatus: successStatus, Res: result}
	return output, nil
}
