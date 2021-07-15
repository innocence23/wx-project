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
	user, exists := c.Get("user")
	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := zerror.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	id := user.(*model.User).ID
	goctx := c.Request.Context()
	u, err := h.UserService.Get(goctx, id)
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
	goctx := c.Request.Context()
	err := h.UserService.Signup(goctx, u)
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

type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

func (h *UserHandler) Signin(c *gin.Context) {
	var req signupReq
	if ok := bindData(c, &req); !ok {
		return
	}
	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	ctx := c.Request.Context()
	err := h.UserService.Signup(ctx, u)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// create token pair as strings
	tokens, err := h.TokenService.NewPairFromUser(ctx, u, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())
		c.JSON(zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}

//todo 如何设置失效？ 刷新token？ 	return s.TokenRepository.DeleteUserRefreshTokens(ctx, uid.String())

func (h *UserHandler) Signout(c *gin.Context) {
	user := c.MustGet("user")

	ctx := c.Request.Context()
	if err := h.TokenService.Signout(ctx, user.(*model.User).UID); err != nil {
		c.JSON(zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user signed out successfully!",
	})
}

type detailsReq struct {
	Name   string `json:"name" binding:"omitempty,max=50"`
	Avatar string `json:"avatar" binding:"required"`
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	authUser := c.MustGet("user").(*model.User)
	var req detailsReq
	if ok := bindData(c, &req); !ok {
		return
	}
	u := &model.User{
		ID:     authUser.ID,
		Name:   req.Name,
		Avatar: req.Avatar,
	}

	ctx := c.Request.Context()
	err := h.UserService.UpdateDetail(ctx, u)

	if err != nil {
		log.Printf("Failed to update user: %v\n", err.Error())

		c.JSON(zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
