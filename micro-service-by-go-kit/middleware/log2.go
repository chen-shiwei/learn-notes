package middleware

import (
	"chen/micro-service/service"
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingMiddleware2 struct {
	Logger log.Logger
	Next   service.StringServiceInterface
}

func (lm2 LoggingMiddleware2) Uppercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		lm2.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = lm2.Next.Uppercase(ctx, s)
	return
}

func (lm2 LoggingMiddleware2) Count(ctx context.Context, s string) (n int) {
	defer func(begin time.Time) {
		lm2.Logger.Log(
			"method", "uppercase",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())
	n = lm2.Next.Count(ctx, s)
	return
}
