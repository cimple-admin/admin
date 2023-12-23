package validate

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
)

var Validate *validator.Validate

func Init() {
	en := en.New()
	cn := zh_Hans.New()

	uni := ut.New(cn, en)
	trans, _ := uni.GetTranslator("zh_Hans")
	Validate = validator.New()
	zh.RegisterDefaultTranslations(Validate, trans)
	Validate.RegisterValidation("password", validatePassword)
}
