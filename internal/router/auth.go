package router

import (
	apiAuth "github.com/cimple-admin/admin/internal/api/auth"
	"github.com/gofiber/fiber/v2"
)

func auth(app *fiber.App) {
	app.Post("/login", apiAuth.Login)
	app.Post("/register", apiAuth.Register)
}
