package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

type user struct {
	Email  string `json:"email" xml:"email" form:"email" validate:"required|email" label:"邮箱"`
	Pass   string `json:"pass" xml:"pass" form:"pass" validate:"required|minLength:5|maxLength:20" label:"密码"`
	RePass string `json:"repass" xml:"repass" form:"repass" validate:"required|eq_field:Pass" label:"确认密码"`
}

func Register(ctx *fiber.Ctx) error {
	u := new(user)
	if err := ctx.BodyParser(u); err != nil {
		return ctx.JSON(fiber.Map{
			"status": -1,
			"error":  err.Error(),
		})
	}

	v := validate.Struct(u)
	if !v.Validate() {
		ctx.JSON(fiber.Map{
			"status": -2,
			"email":  v.Errors,
		})
	}

	return nil
}
