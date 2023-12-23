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
	lis := viper.GetString("LISTEN")
	fmt.Println("l:" + lis)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(lis)
}
