package main

import (
	"chen/micro-service/middleware"
	"chen/micro-service/service"
	"context"
	"encoding/json"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func main() {
	var svc service.StringServiceInterface
	svc = service.StringService{}

	logger := log.NewLogfmtLogger(os.Stderr)

	// lm // var uppercase endpoint.Endpoint
	// uppercase = service.MakeUppercaseEndpoint(svc)
	// uppercase = middleware.LoggineMiddleware(log.With(logger, "method", "uppercase"))(uppercase)

	// var count endpoint.Endpoint
	// count = service.MakeCountEndpoint(svc)
	// count = middleware.LoggineMiddleware(log.With(logger, "method", "count"))(count)

	svc = middleware.LoggingMiddleware2{logger, svc}

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here
	svc = middleware.InstrumentingMiddle{RequestCount: requestCount, RequestLatency: requestLatency, CountResult: countResult, Next: svc}

	uppercaseHandler := httptransport.NewServer(
		service.MakeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		service.MakeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", stdprometheus.Handler())
	logger.Log("msg", "HTTP", "addr", "8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))

}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// var request
	var request service.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request service.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
