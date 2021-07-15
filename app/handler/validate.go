package handler

import (
	"fmt"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// used to help extract validation errors
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func bindData(ctx *gin.Context, req interface{}) bool {
	if err := ctx.Bind(req); err != nil {
		fmt.Println(err)
		fmt.Printf("%#v, %T\n", err, err)
		fmt.Println("-----")
		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []invalidArgument
			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field,
					err.Value.(string),
					err.Tag,
					err.Param,
				})
			}
			err := zerror.NewBadRequest("参数非法")
			Fail(ctx, err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			fmt.Println(err)

			return false
		}

		fallBack := zerror.NewInternal()
		Fail(ctx, fallBack.Status(), gin.H{"error": fallBack})
		return false
	}
	return true
}
