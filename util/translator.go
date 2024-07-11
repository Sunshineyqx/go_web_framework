package util

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

var trans *ut.Translator

// Log 返回日志对象
func Translator() ut.Translator {
	if trans == nil {
		InitTrans(os.Getenv("TRANSLATE_LANGUAGE"))
	}
	return *trans
}

func InitTrans(locale string) (err error) {
	// 修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		enT := en.New()
		// 第一个参数是备用的语言，后面是应该支持的语言
		uni := ut.New(enT, zhT, enT)

		if trans == nil {
			trans = new(ut.Translator)
		}
		*trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) error", locale)
		}

		switch locale {
		case "en":
			en_trans.RegisterDefaultTranslations(v, *trans)
		case "zh":
			zh_trans.RegisterDefaultTranslations(v, *trans)
		default:
			zh_trans.RegisterDefaultTranslations(v, *trans)
		}
		return
	}
	return
}
