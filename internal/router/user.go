package router

import (
	authMiddleware "github.com/cimple-admin/admin/internal/middleware/auth"
	"github.com/gofiber/fiber/v2"
)

func user(app *fiber.App) {
	userGroup := app.Group("/user", authMiddleware.Auth())
	userGroup.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"info": "ok",
		})
	})
}
