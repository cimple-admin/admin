package auth

import (
	"github.com/cimple-admin/admin/internal/model"
	"github.com/cimple-admin/admin/internal/response/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Email    string `json:"email" xml:"email" form:"email" validate:"required|email" label:"邮箱"`
	Password string `json:"password" xml:"password" form:"password" validate:"required" label:"密码"`
}

func Login(ctx *fiber.Ctx) error {
	u := new(login)
	if err := ctx.BodyParser(u); err != nil {
		return json.Fail(ctx, -1, err.Error())
	}

	v := validate.Struct(u)
	if !v.Validate() {
		return json.Fail(ctx, -2, v.Errors.One())
	}
	var user model.User
	model.DB.Where("email = ?", u.Email).First(&user)
	if user.ID == 0 {
		return json.Fail(ctx, -3, "用户不存在")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return json.Fail(ctx, -4, "密码不正确")
	}

	encryptedToken, err := generateToken(user.ID)
	if err != nil {
		return json.Fail(ctx, -5, "generateTokenError")
	}

	return json.SuccessData(ctx, "登录成功", encryptedToken)
}
