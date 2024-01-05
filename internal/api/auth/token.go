package auth

import (
	"strconv"
	"time"

	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/spf13/viper"
)

func generateToken(uid uint) (string, error) {
	return pasetoware.CreateToken([]byte(viper.GetString("PASETOKEY")), strconv.FormatUint(uint64(uid), 10), 12*time.Hour, pasetoware.PurposeLocal)
}
