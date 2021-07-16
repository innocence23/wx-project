package handler

import (
	"log"
	"wx/app/component"
	"wx/app/dto"
	"wx/app/handler/commonhandler"
	"wx/app/handler/middleware"
	"wx/app/iface"
	"wx/app/model"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type userHandler struct {
	UserService iface.UserService
}

func NewUserHandler(s iface.UserService) *userHandler {
	return &userHandler{
		UserService: s,
	}
}

func (h *userHandler) Router(router *gin.RouterGroup) {
	//xrouter := router.Group("/x")
	router.POST("/signup", h.signup)
	router.POST("/signin", h.signin)

	router.Use(middleware.JWTAuthMiddleware())
	router.POST("/signout", h.signout)
	router.PUT("/user", h.updateUser)
	router.GET("/user/info", h.me)
	router.PUT("/user/enable", h.enableUser)
	router.PUT("/user/disable", h.disableUser)
}

func (h *userHandler) me(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		log.Printf("上下文中获取不到用户: %v\n", ctx)
		e := zerror.NewInternal()
		commonhandler.Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}

	id := user.(*dto.UserJWT).ID
	goctx := ctx.Request.Context()
	result, err := h.UserService.Get(goctx, id)
	if err != nil {
		log.Printf("无法找到用户: %v %#v\n%v", id, user, err)
		e := zerror.NewNotFound("user", cast.ToString(id))
		commonhandler.Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}
	commonhandler.Success(ctx, result)
}

func (h *userHandler) signup(ctx *gin.Context) {
	var req dto.SignupReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
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
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}

	uj := &dto.UserJWT{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		Avatar: u.Avatar,
	}
	tokens, err := component.GenerateToken(uj)
	if err != nil {
		log.Printf("token生成失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"tokens": tokens,
	})
}

func (h *userHandler) signin(ctx *gin.Context) {
	var req dto.SigninReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	goctx := ctx.Request.Context()
	err := h.UserService.Signin(goctx, u)
	if err != nil {
		log.Printf("登陆失败 user: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}

	uj := &dto.UserJWT{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		Avatar: u.Avatar,
	}
	tokens, err := component.GenerateToken(uj)
	if err != nil {
		log.Printf("token生成失败 user: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"tokens": tokens,
	})
}

//todo 如何让客户端的token失效
func (h *userHandler) signout(ctx *gin.Context) {
	authUser := ctx.MustGet("user").(*dto.UserJWT)
	_ = authUser
	commonhandler.Success(ctx, gin.H{
		"message": "退出成功",
	})
}

func (h *userHandler) updateUser(ctx *gin.Context) {
	var req dto.DetailsReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}

	authUser := ctx.MustGet("user").(*dto.UserJWT)
	u := &model.User{
		ID:     authUser.ID,
		Name:   req.Name,
		Avatar: req.Avatar,
	}
	goctx := ctx.Request.Context()
	err := h.UserService.UpdateDetail(goctx, u)
	if err != nil {
		log.Printf("更新用户失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}

	commonhandler.Success(ctx, gin.H{
		"message": "操作成功",
	})
}

func (h *userHandler) disableUser(ctx *gin.Context) {
	authUser := ctx.MustGet("user").(*dto.UserJWT)
	goctx := ctx.Request.Context()
	err := h.UserService.DisableUser(goctx, authUser.ID)
	if err != nil {
		log.Printf("用户禁用失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"message": "操作成功",
	})
}

func (h *userHandler) enableUser(ctx *gin.Context) {
	authUser := ctx.MustGet("user").(*dto.UserJWT)
	goctx := ctx.Request.Context()
	err := h.UserService.EnableUser(goctx, authUser.ID)
	if err != nil {
		log.Printf("用户启用失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"message": "操作成功",
	})
}
