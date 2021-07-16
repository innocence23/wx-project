package commonhandler

import (
	"reflect"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)

type invalidArgument struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
	Tag   string      `json:"tag"`
	Param string      `json:"param"`
	Msg   string      `json:"message"`
}

var Trans ut.Translator

func init() {
	uni := ut.New(zh.New())
	Trans, _ = uni.GetTranslator("zh")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = zh2.RegisterDefaultTranslations(v, Trans) //注册翻译器
		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})
	}
}

func BindData(ctx *gin.Context, req interface{}) bool {
	if err := ctx.Bind(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []invalidArgument
			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value(),
					err.Tag(),
					err.Param(),
					err.Translate(Trans),
				})
			}
			err := zerror.NewBadRequest("参数非法")
			Fail(ctx, err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			return false
		}
		fallBack := zerror.NewInternal()
		Fail(ctx, fallBack.Status(), gin.H{"error": fallBack})
		return false
	}
	return true
}
