package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	l := viper.GetString("Listen")
	fmt.Println("l:" + l)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(l)
}
