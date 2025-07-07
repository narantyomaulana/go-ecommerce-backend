package handler

import (
	"context"
	"fmt"

	"github.com/narantyomaulana/go-grpc-ercommerce-be/internal/utils"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/pb/service"
)

type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

func (sh *serviceHandler) HelloWord(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	// Validate the request using protovalidate
	validationErrors, err := utils.CheckValidation(request)

	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &service.HelloWorldResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	return &service.HelloWorldResponse{
		Message: fmt.Sprintf("hello %s, welcome to our service", request.Name),
		Base:    utils.SuccessResponse("Success"),
	}, nil
}

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
