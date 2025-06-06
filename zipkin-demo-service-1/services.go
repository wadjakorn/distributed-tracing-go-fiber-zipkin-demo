package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

func FetchOrders(c *fiber.Ctx) error {
	span, ctx := GetTracer().StartSpanFromContext(c.Context(), "FetchOrders")
	defer span.Finish()
	httpReq, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:5000/api/v1/orders", nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("request-id", c.Get("request-id"))
	resp, err := Client.DoRequest(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	var orders []Order
	if err := json.Unmarshal(b, &orders); err != nil {
		return fmt.Errorf("failed to unmarshal orders: %w", err)
	}
	return c.Status(fiber.StatusOK).JSON(orders)
}

func PlaceOrder(c *fiber.Ctx) error {
	span, ctx := GetTracer().StartSpanFromContext(c.Context(), "PlaceOrder")
	defer span.Finish()
	reqBody := c.Body()
	httpReq, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:5000/api/v1/order", bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("request-id", c.Get("request-id"))
	resp, err := Client.DoRequest(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create order, status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	var order Order
	if err := json.Unmarshal(b, &order); err != nil {
		return fmt.Errorf("failed to unmarshal order: %w", err)
	}
	return c.Status(fiber.StatusCreated).JSON(order)
}
