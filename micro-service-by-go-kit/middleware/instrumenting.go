package middleware

import (
	"chen/micro-service/service"
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddle struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           service.StringServiceInterface
}

func (im InstrumentingMiddle) Uppercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "uppercase", "error", fmt.Sprint(err != nil)}
		im.RequestCount.With(lvs...).Add(1)
		im.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = im.Next.Uppercase(ctx, s)
	return
}

func (im InstrumentingMiddle) Count(ctx context.Context, s string) (n int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "count", "error", "false"}
		im.RequestCount.With(lvs...).Add(1)
		im.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		im.CountResult.Observe(float64(n))
	}(time.Now())
	n = im.Next.Count(ctx, s)
	return
}
