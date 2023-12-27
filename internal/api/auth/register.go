package auth

import (
	"github.com/cimple-admin/admin/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Email           string `json:"email" xml:"email" form:"email" validate:"required|email" label:"邮箱"`
	Password        string `json:"password" xml:"password" form:"password" validate:"required|minLength:5|maxLength:20|password" label:"密码"`
	ConfirmPassword string `json:"confirm_password" xml:"confirm_password" form:"confirm_password" validate:"required|eq_field:Password" label:"确认密码"`
}

func Register(ctx *fiber.Ctx) error {
	u := new(user)
	if err := ctx.BodyParser(u); err != nil {
		return ctx.JSON(fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}

	v := validate.Struct(u)
	if !v.Validate() {
		ctx.JSON(fiber.Map{
			"status":  -2,
			"message": v.Errors.One(),
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		ctx.JSON(fiber.Map{
			"status":  -3,
			"message": err,
		})
	}
	user := model.User{Email: u.Email, Password: string(password), Name: u.Email}
	result := model.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(fiber.Map{
			"status":  -4,
			"message": result.Error,
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  0,
		"message": "注册成功",
	})
}
