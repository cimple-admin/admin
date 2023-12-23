package router

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	auth(app)
	common(app)
}
