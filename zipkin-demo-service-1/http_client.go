package main

import (
	"net/http"

	"github.com/openzipkin/zipkin-go"
	zipkin_http_middleware "github.com/openzipkin/zipkin-go/middleware/http"
)

var Client IClient

type IClient interface {
	DoRequest(req *http.Request) (*http.Response, error)
}

type HttpClient struct {
	Tracer           *zipkin.Tracer
	zipkinHttpClient *zipkin_http_middleware.Client
}

func NewHttpClient() {
	zipkinHttpClient, _ := zipkin_http_middleware.NewClient(GetTracer(), zipkin_http_middleware.ClientTrace(true))
	Client = &HttpClient{
		Tracer:           GetTracer(),
		zipkinHttpClient: zipkinHttpClient,
	}
}

func (hc *HttpClient) DoRequest(req *http.Request) (*http.Response, error) {
	resp, err := hc.zipkinHttpClient.DoWithAppSpan(req, "http_client")
	if err != nil {
		return nil, err
	}
	return resp, nil
}
