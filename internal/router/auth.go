package router

import (
	"github.com/cimple-admin/admin/internal/api/auth"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuth(app *fiber.App) {
	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)
}
