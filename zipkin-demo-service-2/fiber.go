package main

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/adaptor/v2"
	zipkin_http_middleware "github.com/openzipkin/zipkin-go/middleware/http"
)

var (
	AppName = "zipkin-demo-service-2"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func FiberWebServer() (*fiber.App, context.Context) {
	app := fiber.New(fiber.Config{
		AppName:        AppName,
		ReadBufferSize: 60000, // default 60KB
		Immutable:      true,
	})

	app.Use(adaptor.HTTPMiddleware(zipkin_http_middleware.NewServerMiddleware(
		GetTracer(),
		zipkin_http_middleware.SpanName("http_server"),
	)))

	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
				Message: err.Error(),
			})
		}
		return err
	})

	return app, context.Background()
}
