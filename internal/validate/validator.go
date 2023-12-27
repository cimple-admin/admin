package validate

import (
	"github.com/gookit/validate"
	"github.com/gookit/validate/locales/zhcn"
)

func Init() {
	zhcn.RegisterGlobal()
	validate.AddGlobalMessages(map[string]string{
		"eqField": "{field} 字段值必须和 %s 字段值相同",
	})
}
