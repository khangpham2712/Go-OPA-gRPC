package multiplication

import (
	"context"
	"dummy/proto"
	"log"
	"time"
)

type MultiplicationServer struct {
	proto.UnimplementedMultiplicationServer
}

func (server *MultiplicationServer) Multiply(c context.Context, input *proto.Input) (*proto.Output, error) {
	//this code block is used to log the method execution time
	start := time.Now()
	defer func(start time.Time) {
		duration := time.Since(start)
		log.Println("Mul service duration:", duration)
	}(start)
	//
	a, b := input.GetA(), input.GetB()
	result := a * b
	var successStatus int32 = 1
	var output *proto.Output = &proto.Output{SuccessStatus: successStatus, Res: result}
	// time.Sleep(200 * time.Millisecond)
	return output, nil
}
