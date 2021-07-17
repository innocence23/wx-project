package handler

import (
	"log"
	"wx/app/component"
	"wx/app/dto"
	"wx/app/handler/commonhandler"
	"wx/app/iface"
	"wx/app/model"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type permissionHandler struct {
	PermissionService iface.PermissionService
}

func NewPermissionHandler(s iface.PermissionService) *permissionHandler {
	return &permissionHandler{
		PermissionService: s,
	}
}

func (h *permissionHandler) Router(router *gin.RouterGroup) {
	grouter := router.Group("/permission")
	grouter.GET("/info", h.show)
	grouter.POST("/list", h.list) //参数多，改为post方便
	grouter.POST("", h.create)
	grouter.PUT("", h.update)
	grouter.PUT("/enable", h.enable)
	grouter.PUT("/disable", h.disable)
	grouter.PUT("/generate", h.generate)
}

func (h *permissionHandler) show(ctx *gin.Context) {
	var req dto.PermissionIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	result, err := h.PermissionService.Find(goctx, id)
	if err != nil {
		log.Printf("信息不存在: %v \n%v", id, err)
		e := zerror.NewNotFound("permission", cast.ToString(id))
		commonhandler.Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}
	commonhandler.Success(ctx, result)
}

func (h *permissionHandler) list(ctx *gin.Context) {
	var req dto.PermissionSearchReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	where := req
	goctx := ctx.Request.Context()
	list, err := h.PermissionService.List(goctx, where)
	if err != nil {
		log.Printf("无法找到数据: %#v\n%v", where, err)
		e := zerror.NewInternal()
		commonhandler.Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}
	commonhandler.Success(ctx, list)
}

func (h *permissionHandler) create(ctx *gin.Context) {
	var req dto.PermissionCreateReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	data := &model.Permission{
		Name:   req.Name,
		Group:  req.Group,
		Url:    req.Url,
		Method: req.Method,
	}
	goctx := ctx.Request.Context()
	result, err := h.PermissionService.Create(goctx, data)
	if err != nil {
		log.Printf("数据创建失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, result)
}

func (h *permissionHandler) update(ctx *gin.Context) {
	var req dto.PermissionUpdateReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	data := &model.Permission{
		Id:     req.ID,
		Name:   req.Name,
		Group:  req.Group,
		Url:    req.Url,
		Method: req.Method,
	}
	goctx := ctx.Request.Context()
	err := h.PermissionService.Update(goctx, data)
	if err != nil {
		log.Printf("数据更新失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"message": "操作成功",
	})
}

func (h *permissionHandler) disable(ctx *gin.Context) {
	var req dto.PermissionIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	if err := h.PermissionService.Disable(goctx, id); err != nil {
		log.Printf("数据禁用失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"message": "操作成功",
	})
}

func (h *permissionHandler) enable(ctx *gin.Context) {
	var req dto.PermissionIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	if err := h.PermissionService.Enable(goctx, id); err != nil {
		log.Printf("数据启用失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"message": "操作成功",
	})
}

func (h *permissionHandler) generate(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		log.Printf("上下文中获取不到用户: %v\n", ctx)
		e := zerror.NewInternal()
		commonhandler.Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}
	if user.(*dto.UserJWT).Email != "admin@admin.com" {
		e := zerror.NewAuthorization("没有权限")
		commonhandler.Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
	}

	routers := component.GetAllRoutes()
	goctx := ctx.Request.Context()
	if err := h.PermissionService.AutoGenerate(goctx, routers); err != nil {
		log.Printf("自动化失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"message": "操作成功",
	})
}
