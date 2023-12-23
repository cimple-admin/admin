package validate

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func validatePassword(fl validator.FieldLevel) bool {
	hasNumber := false
	hasAlpha := false
	hasAlphaUpper := false
	hasSymbol := false

	pass := fl.Field().String()

	for _, c := range pass {
		switch {
		case unicode.IsLetter(c):
			hasAlpha = true
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasAlphaUpper = true
		case unicode.IsSymbol(c):
			hasSymbol = true
		}
	}

	return hasAlpha && hasAlphaUpper && hasNumber && hasSymbol
}
