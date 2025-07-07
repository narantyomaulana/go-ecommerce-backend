package utils

import (
	"errors"

	"buf.build/go/protovalidate"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/pb/common"
	"google.golang.org/protobuf/proto"
)

func CheckValidation(req proto.Message) ([]*common.ValidationError, error) {
	var validationErrorResponse []*common.ValidationError = make([]*common.ValidationError, 0)

	if err := protovalidate.Validate(req); err != nil {
		var validationError *protovalidate.ValidationError

		// Check if the error is a validation error
		if errors.As(err, &validationError) {
			for _, violation := range validationError.Violations {
				validationErrorResponse = append(validationErrorResponse, &common.ValidationError{
					Field:   *violation.Proto.Field.Elements[0].FieldName,
					Message: *violation.Proto.Message,
				})
			}
			return validationErrorResponse, nil
		}
		return nil, err
	}

	return nil, nil
}
