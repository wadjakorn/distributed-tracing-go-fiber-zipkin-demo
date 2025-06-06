package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

type ReqCreateOrder struct {
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

func GetOrders(c *fiber.Ctx) error {
	span, ctx := GetTracer().StartSpanFromContext(c.Context(), "GetOrders")
	defer span.Finish()
	mockOrders := []Order{
		{ID: "1", Amount: 100.0, Status: "Pending"},
		{ID: "2", Amount: 200.0, Status: "Completed"},
		{ID: "3", Amount: 300.0, Status: "Shipped"},
		{ID: "4", Amount: 400.0, Status: "Cancelled"},
	}
	queryOrders(ctx)
	return c.Status(fiber.StatusOK).JSON(mockOrders)
}

func queryOrders(ctx context.Context) {
	span, _ := GetTracer().StartSpanFromContext(ctx, "queryOrders")
	defer span.Finish()

	span.Tag("db.collection", "orders")
	span.Tag("db.operation", "select")

	// sleep to simulate processing time
	time.Sleep(30 * time.Millisecond)

	span.Annotate(time.Now(), "Preparing query for orders")

	// Simulate a database operation
	time.Sleep(100 * time.Millisecond)

	span.Annotate(time.Now(), "Orders fetched from DB")
}

func CreateOrder(c *fiber.Ctx) error {
	span, ctx := GetTracer().StartSpanFromContext(c.Context(), "CreateOrder")
	defer span.Finish()
	var req ReqCreateOrder
	if err := c.BodyParser(&req); err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}
	if req.Amount <= 0 {
		return fmt.Errorf("invalid amount: %f", req.Amount)
	}
	if req.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}
	queryOrders(ctx)
	createOrderInDB(ctx)
	resp := Order{
		ID:     "new-order-id",
		Amount: req.Amount,
		Status: req.Status,
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func createOrderInDB(ctx context.Context) {
	span, _ := GetTracer().StartSpanFromContext(ctx, "createOrderInDB")
	defer span.Finish()
	span.Tag("db.collection", "orders")
	span.Tag("db.operation", "insert")

	// Simulate a database operation
	time.Sleep(100 * time.Millisecond)

	span.Annotate(time.Now(), "Order created in DB")

	// Simulate a database operation
	time.Sleep(100 * time.Millisecond)

	span.Annotate(time.Now(), "New order indexed in search service")
}
