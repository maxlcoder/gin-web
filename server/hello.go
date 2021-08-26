package server

import (
	"context"
	pb "github.com/maxlcoder/gin-web/proto"
)

type helloService struct {

}

func NewHelloService() *helloService {
	return &helloService{}
}

func (h helloService) SayHelloWorld(c context.Context, r *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error)  {
	return &pb.HelloWorldResponse{
		Message: "test",
	}, nil
}
