package utils

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func ValidateInit() {
	uni := ut.New(zh.New())
	Trans, _ = uni.GetTranslator("zh")

	if v, ok := binding.Validator.(*validator.Validate); ok {
		//注册翻译器
		_ = zh2.RegisterDefaultTranslations(v, Trans)
		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})
	}
}

//处理gin错误转换中文
func ValidateParams(err error) string {
	msg := "参数解析错误"
	if errs, ok := err.(validator.ValidationErrors); ok {
		//随机取一个值返回，方便前端显示错误
		for _, val := range errs.Translate(Trans) {
			msg = val
			break
		}
	} else {
		msg = err.Error()
	}
	return msg
}
