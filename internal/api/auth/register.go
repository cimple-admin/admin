package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	Email  string `json:"email" xml:"email" form:"email" validate:"required,min=5,max=20,email"`
	Pass   string `json:"pass" xml:"pass" form:"pass" validate:"required,min=5,max=20,alphanum"`
	RePass string `json:"repass" xml:"repass" form:"repass" validate:"required,eqfield=Pass"`
}

func Register(ctx *fiber.Ctx) error {
	u := new(user)
	if err := ctx.BodyParser(u); err != nil {
		return ctx.JSON(fiber.Map{
			"status": -1,
			"error":  err.Error(),
		})
	}
	if errs := validator.New().Struct(u); errs != nil {
		return ctx.JSON(fiber.Map{
			"status": -2,
			"error":  errs.Error(),
		})
	}

	return nil
}
