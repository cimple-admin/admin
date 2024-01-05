package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/cimple-admin/admin/internal/model"
	"github.com/cimple-admin/admin/internal/response/json"
	"github.com/go-sql-driver/mysql"
	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/spf13/viper"
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
		return json.Fail(ctx, -1, err.Error())
	}

	v := validate.Struct(u)
	if !v.Validate() {
		return json.Fail(ctx, -2, v.Errors.One())
	}

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return json.Fail(ctx, -3, err.Error())
	}
	user := model.User{Email: u.Email, Password: string(password), Name: u.Email}
	result := model.DB.Create(&user)

	if result.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			return json.Fail(ctx, -4, "用户名已存在")
		} else {
			return json.Fail(ctx, -4, result.Error.Error())
		}
	}

	// 注册成功后，生成 token 并返回
	encryptedToken, err := pasetoware.CreateToken([]byte(viper.GetString("PASETOKEY")), strconv.FormatUint(uint64(user.ID), 10), 12*time.Hour, pasetoware.PurposeLocal)
	if err != nil {
		return json.Fail(ctx, -5, "generateTokenError")
	}

	return json.SuccessData(ctx, "注册成功", encryptedToken)
}
