package auth

import (
	"encoding/json"
	"fmt"

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
			err := json.Unmarshal(decrypted, &payload)
			if err == nil {
				uid := payload["data"]
				fmt.Println(uid)
			}
			return payload, err
		},
	})
}
