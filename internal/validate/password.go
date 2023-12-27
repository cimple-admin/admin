package validate

import (
	"unicode"
)

func password(pass string) bool {
	hasNumber := false
	hasAlpha := false
	hasAlphaUpper := false

	for _, c := range pass {
		switch {
		case unicode.IsUpper(c):
			hasAlphaUpper = true
		case unicode.IsLetter(c):
			hasAlpha = true
		case unicode.IsNumber(c):
			hasNumber = true
		}
	}

	return hasAlpha && hasAlphaUpper && hasNumber
}
