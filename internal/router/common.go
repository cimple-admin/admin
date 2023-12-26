package router

import (
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func common(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fiberi18n.MustLocalize(c, &i18n.LocalizeConfig{
			MessageID: "welcome",
		}))
	})
}
