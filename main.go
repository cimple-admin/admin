package main

import (
	"github.com/cimple-admin/admin/internal/config"
	"github.com/cimple-admin/admin/internal/model"
	"github.com/cimple-admin/admin/internal/router"
	"github.com/cimple-admin/admin/internal/validate"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

func init() {
	config.Init()
	model.Init()
	validate.Init()
}

func main() {
	lis := viper.GetString("LISTEN")

	app := fiber.New()

	app.Use(
		fiberi18n.New(&fiberi18n.Config{
			RootPath:        "./lang",
			AcceptLanguages: []language.Tag{language.Chinese},
			DefaultLanguage: language.Chinese,
		}),
	)

	router.Register(app)

	app.Listen(lis)
}
