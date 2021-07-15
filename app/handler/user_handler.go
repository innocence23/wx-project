package handler

import (
	"log"
	"net/http"
	"wx/app/component"
	"wx/app/model"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type UserHandler struct {
	UserService model.UserService
}

func (h *UserHandler) Me(ctx *gin.Context) {
	var id int64 = 10
	u, err := h.UserService.Get(ctx, id)
	if err != nil {
		log.Printf("无法找到用户: %v\n%v", id, err)
		e := zerror.NewNotFound("user", cast.ToString(id))
		fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}
	success(ctx, gin.H{
		"user": u,
	})
}

type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signup handler
func (h *UserHandler) Signup(ctx *gin.Context) {
	var req signupReq
	if !bindData(ctx, &req) {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := h.UserService.Signup(ctx, u)
	if err != nil {
		log.Printf("注册失败: %v\n", err.Error())
		ctx.JSON(zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// create token pair as strings
	tokens, err := component.GenerateToken(u)
	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())
		ctx.JSON(zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}

// Signin handler
func (h *UserHandler) Signin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's signin",
	})
}

// Signout handler
func (h *UserHandler) Signout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's signout",
	})
}

// Tokens handler
func (h *UserHandler) Tokens(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's tokens",
	})
}

// Image handler
func (h *UserHandler) Image(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's image",
	})
}

// DeleteImage handler
func (h *UserHandler) DeleteImage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's deleteImage",
	})
}

// Details handler
func (h *UserHandler) Details(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
