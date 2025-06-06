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
	filterOrders(ctx)
	return c.Status(fiber.StatusOK).JSON(mockOrders)
}

func filterOrders(ctx context.Context) {
	span, _ := GetTracer().StartSpanFromContext(ctx, "filterOrders")
	defer span.Finish()

	// sleep to simulate processing time
	time.Sleep(30 * time.Millisecond)
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

	// sleep to simulate processing time
	time.Sleep(100 * time.Millisecond)
}
