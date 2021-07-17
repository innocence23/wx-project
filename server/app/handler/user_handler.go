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
	router = router.Group("/uc")
	router.POST("/signup", h.signup)
	router.POST("/signin", h.signin)
	router.Use(middleware.JWTAuthMiddleware())
	router.POST("/signout", h.signout)
	router.GET("/me", h.me)
	router.PUT("/update", h.update)

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
	result := user.(*dto.UserJWT)
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

	uj := h.genUserJwt(ctx, u)
	token, err := component.GenerateToken(uj)
	if err != nil {
		log.Printf("token生成失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"token": token,
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

	uj := h.genUserJwt(ctx, u)
	token, err := component.GenerateToken(uj)
	if err != nil {
		log.Printf("token生成失败 user: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"token": token,
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

func (h *userHandler) update(ctx *gin.Context) {
	var req dto.UserUpdateReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}

	authAdmin := ctx.MustGet("user").(*dto.UserJWT)
	u := &model.User{
		ID:      authAdmin.ID,
		Account: req.Account,
		Avatar:  req.Avatar,
	}
	goctx := ctx.Request.Context()
	err := h.UserService.Update(goctx, u)
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

//组装jwt数据
func (h *userHandler) genUserJwt(ctx *gin.Context, u *model.User) *dto.UserJWT {
	uj := &dto.UserJWT{
		ID:      u.ID,
		Account: u.Account,
		Email:   u.Email,
		Avatar:  u.Avatar,
	}
	id := cast.ToString(uj.ID)
	uj.Roles, _ = component.GetRolesForUser(id)
	uj.Permissions = component.GetPermissionsForUser(id)
	var tmp []int64
	for _, v := range uj.Roles {
		tmp = append(tmp, cast.ToInt64(v))
	}
	goctx := ctx.Request.Context()
	uj.Menus = h.UserService.GetMenus(goctx, tmp, u.Email)

	return uj
}
