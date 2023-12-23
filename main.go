package main

import (
	"fmt"

	"github.com/cimple-admin/admin/internal/config"
	"github.com/cimple-admin/admin/internal/router"
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

	router.Register(app)

	app.Listen(lis)
}
