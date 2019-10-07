package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeUppercaseEndpoint(svc StringServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(ctx, req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}
		return UppercaseResponse{v, ""}, nil
	}
}

func MakeCountEndpoint(svc StringServiceInterface) endpoint.Endpoint {
	return func(cxt context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(cxt, req.S)
		return CountResponse{v}, nil
	}
}
