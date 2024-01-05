package request

import (
	"github.com/cimple-admin/admin/internal/model"
	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
)

func GetLoginUser(ctx *fiber.Ctx) model.User {
	return ctx.Locals(pasetoware.DefaultContextKey).(model.User)
}
