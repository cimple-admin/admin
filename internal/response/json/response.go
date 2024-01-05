package json

import "github.com/gofiber/fiber/v2"

func Fail(ctx *fiber.Ctx, status int, message string) error {
	return ctx.JSON(fiber.Map{
		"status":  status,
		"message": message,
	})
}

func FailData(ctx *fiber.Ctx, status int, message string, data interface{}) error {
	return ctx.JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func Success(ctx *fiber.Ctx, message string) error {
	return Fail(ctx, 0, message)
}

func SuccessData(ctx *fiber.Ctx, message string, data interface{}) error {
	return FailData(ctx, 0, message, data)
}
