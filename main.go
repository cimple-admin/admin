package main

import (
	"fmt"

	"github.com/cimple-admin/admin/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	config.Init()
}

func main() {
	l := viper.GetString("Listen")
	fmt.Println("l:" + l)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(l)
}
