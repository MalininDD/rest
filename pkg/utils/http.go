package utils

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Read request body and validate
func ReadRequest(c *fiber.Ctx, request interface{}) error {
	if err := c.BodyParser(request); err != nil {
		return err
	}

	return validate.StructCtx(c.Context(), request)
}

//func getRequestID(c *fiber.Ctx) string {
//	return c.Get(fiber.HeaderXRequestID)
//}

//type ReqIDCtxKey struct{}

func GetRequestCtx(c *fiber.Ctx) context.Context {
	requestID := uuid.New().String()
	c.Set(fiber.HeaderXRequestID, requestID)
	return context.WithValue(c.Context(), fiber.HeaderXRequestID, requestID)
}
