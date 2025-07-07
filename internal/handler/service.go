package handler

import (
	"context"
	"fmt"

	"github.com/narantyomaulana/go-grpc-ercommerce-be/pb/service"
)

type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

func (sh *serviceHandler) HelloWord(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	return &service.HelloWorldResponse{
		Message: fmt.Sprintf("hello %s, welcome to our service", request.Name),
	}, nil
}

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
