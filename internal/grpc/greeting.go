package grpc

import (
	"context"
	"fmt"
	proto "github.com/ElladanTasartir/buffy-grpc/gen/go/greeting/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreetingService struct{}

func (gs GreetingService) Greet(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	if req.Msg == nil {
		err := status.New(codes.InvalidArgument, "Message cannot be empty").Err()
		return nil, err
	}

	helloMsg := fmt.Sprintf("%s, %s", req.Msg.Greeting.String(), req.Msg.Name)
	res := &proto.GreetResponse{
		Resp: helloMsg,
	}

	return res, nil
}
