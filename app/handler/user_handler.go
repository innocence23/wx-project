package handler

import (
	"log"
	"wx/app/component"
	"wx/app/dto"
	"wx/app/model"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type userHandler struct {
	UserService model.UserService
}

// 注册路由
func (h *userHandler) Router(router *gin.RouterGroup) {
	//xrouter := router.Group("/x")
	router.POST("/signup", h.signup)
	router.POST("/signin", h.signin)
	router.POST("/signout", h.signout)
	router.PUT("/user", h.updateUser)
	router.GET("/user", h.me)
}

func NewUserHandler(us model.UserService) *userHandler {
	return &userHandler{
		UserService: us,
	}
}

func (h *userHandler) me(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		log.Printf("上下文中获取不到用户: %v\n", ctx)
		e := zerror.NewInternal()
		Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}

	id := user.(*model.User).ID
	goctx := ctx.Request.Context()
	u, err := h.UserService.Get(goctx, id)
	if err != nil {
		log.Printf("无法找到用户: %v\n%v", id, err)
		e := zerror.NewNotFound("user", cast.ToString(id))
		Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}
	Success(ctx, gin.H{
		"user": u,
	})
}

func (h *userHandler) signup(ctx *gin.Context) {
	var req dto.SignupReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	goctx := ctx.Request.Context()
	err := h.UserService.Signup(goctx, u)
	if err != nil {
		log.Printf("注册失败: %v\n", err.Error())
		Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}

	tokens, err := component.GenerateToken(u)
	if err != nil {
		log.Printf("token生成失败: %v\n", err.Error())
		Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	Success(ctx, gin.H{
		"tokens": tokens,
	})
}

func (h *userHandler) signin(ctx *gin.Context) {
	var req dto.SigninReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	goctx := ctx.Request.Context()
	err := h.UserService.Signup(goctx, u)
	if err != nil {
		log.Printf("登陆失败 user: %v\n", err.Error())
		Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}

	tokens, err := component.GenerateToken(u)
	if err != nil {
		log.Printf("token生成失败 user: %v\n", err.Error())
		Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	Success(ctx, gin.H{
		"tokens": tokens,
	})
}

func (h *userHandler) signout(ctx *gin.Context) {
	authUser := ctx.MustGet("user").(*model.User)
	//todo 如何设置失效？ 刷新token？ 	return s.TokenRepository.DeleteUserRefreshTokens(ctx, uid.String())
	if _, err := component.GenerateToken(authUser); err != nil {
		Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	Success(ctx, gin.H{
		"message": "退出成功",
	})
}

func (h *userHandler) updateUser(ctx *gin.Context) {
	var req dto.DetailsReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	authUser := ctx.MustGet("user").(*model.User)
	u := &model.User{
		ID:     authUser.ID,
		Name:   req.Name,
		Avatar: req.Avatar,
	}
	goctx := ctx.Request.Context()
	err := h.UserService.UpdateDetail(goctx, u)
	if err != nil {
		log.Printf("更新用户失败: %v\n", err.Error())
		Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}

	Success(ctx, gin.H{
		"user": u,
	})
}
