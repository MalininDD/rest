package utils

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func StartFiberTrace(c *fiber.Ctx, spanName string) (context.Context, trace.Span) {
	requestID := uuid.New().String()
	c.Set(fiber.HeaderXRequestID, requestID)
	ctx := context.WithValue(c.Context(), fiber.HeaderXRequestID, requestID)
	ctx, span := otel.Tracer("").Start(ctx, spanName)
	span.SetAttributes(attribute.String(fiber.HeaderXRequestID, requestID))
	if clientID, ok := c.Locals("clientID").(int); ok {
		span.SetAttributes(attribute.Int("clientID", clientID))
	}
	if clientID, ok := c.Locals("APIID").(int); ok {
		span.SetAttributes(attribute.Int("APIID", clientID))
	}
	if userID, ok := c.Locals("userID").(int); ok {
		span.SetAttributes(attribute.Int("userID", userID))
	}
	return ctx, span
}
