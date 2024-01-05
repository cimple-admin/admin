package auth

import (
	"encoding/json"
	"fmt"

	"github.com/cimple-admin/admin/internal/model"
	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Auth() func(*fiber.Ctx) error {
	return pasetoware.New(pasetoware.Config{
		SymmetricKey: []byte(viper.GetString("PASETOKEY")),
		TokenPrefix:  "Bearer",
		Validate: func(decrypted []byte) (interface{}, error) {
			var payload map[string]interface{}
			var user model.User
			err := json.Unmarshal(decrypted, &payload)
			if err == nil {
				uid := payload["data"]
				fmt.Println(uid)
				model.DB.Find(&user, uid)
			}

			return user, err
		},
	})
}
