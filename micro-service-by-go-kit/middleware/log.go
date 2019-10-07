package middleware

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
)

func LoggineMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "celling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
