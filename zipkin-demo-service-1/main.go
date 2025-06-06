package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	Zipkin()
	NewHttpClient()
	app, ctx := FiberWebServer()

	app.Get("/test", func(c *fiber.Ctx) error {
		response := fiber.Map{
			"message":   "Hello, this is a test endpoint!",
			"requestId": c.Locals("request-id"),
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	app.Get("/orders", FetchOrders)
	app.Post("/order", PlaceOrder)

	if err := app.Listen(":4000"); err != nil {
		panic(err)
	}

	signal := waitExitSignal()
	fmt.Printf("Received signal %q, shutting down...", signal.String())

	if err := stopAppServer(ctx, app, 10*time.Second); err != nil {
		fmt.Printf("Shutdown server failed %v", err.Error())
	}
}

func waitExitSignal() os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	return <-quit
}

func stopAppServer(ctx context.Context, app *fiber.App, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return app.ShutdownWithContext(ctx)
}
