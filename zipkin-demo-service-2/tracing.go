package main

import (
	"github.com/openzipkin/zipkin-go"
	zipkin_model "github.com/openzipkin/zipkin-go/model"
	reporterHttp "github.com/openzipkin/zipkin-go/reporter/http"
)

var globalTracer *zipkin.Tracer

func NewTracer(name string, port uint16, endpoint string, samplingRate float64) error {
	reporter := reporterHttp.NewReporter(endpoint)
	localEndpoint := &zipkin_model.Endpoint{ServiceName: name, Port: port}
	sampler, err := zipkin.NewCountingSampler(samplingRate)
	if err != nil {
		return err
	}

	t, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(localEndpoint),
	)
	if err != nil {
		return err
	}

	globalTracer = t

	return nil
}

func GetTracer() *zipkin.Tracer {
	return globalTracer
}
