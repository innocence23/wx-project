package handler

import (
	"log"
	"wx/app/dto"
	"wx/app/handler/commonhandler"
	"wx/app/model"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type roleHandler struct {
	RoleService model.RoleService
}

func NewRoleHandler(s model.RoleService) *roleHandler {
	return &roleHandler{
		RoleService: s,
	}
}

func (h *roleHandler) Router(router *gin.RouterGroup) {
	grouter := router.Group("/role")
	grouter.GET("/info", h.show)
	grouter.GET("/list", h.list)
	grouter.POST("", h.create)
	grouter.PUT("/detail", h.update)
	grouter.PUT("/enable", h.enable)
	grouter.PUT("/disable", h.disable)
}

func (h *roleHandler) show(ctx *gin.Context) {
	var req dto.RoleIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	data, err := h.RoleService.Get(goctx, id)
	if err != nil {
		log.Printf("信息不存在: %v \n%v", id, err)
		e := zerror.NewNotFound("role", cast.ToString(id))
		commonhandler.Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"role": data,
	})
}

func (h *roleHandler) list(ctx *gin.Context) {
	var req dto.RoleSearchReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	where := req
	goctx := ctx.Request.Context()
	list, err := h.RoleService.List(goctx, where)
	if err != nil {
		log.Printf("无法找到数据: %#v\n%v", where, err)
		e := zerror.NewInternal()
		commonhandler.Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"users": list,
	})
}

func (h *roleHandler) create(ctx *gin.Context) {
	var req dto.RoleCreateReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	data := &model.Role{
		Name:        req.Name,
		Description: req.Description,
		MenuIds:     req.MenuIds,
	}
	goctx := ctx.Request.Context()
	result, err := h.RoleService.Create(goctx, data)
	if err != nil {
		log.Printf("数据创建失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"role": result,
	})
}

func (h *roleHandler) update(ctx *gin.Context) {
	var req dto.RoleUpdateReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	data := &model.Role{
		Id:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		MenuIds:     req.MenuIds,
	}
	goctx := ctx.Request.Context()
	err := h.RoleService.Update(goctx, data)
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

func (h *roleHandler) disable(ctx *gin.Context) {
	var req dto.RoleIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	if err := h.RoleService.Disable(goctx, id); err != nil {
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

func (h *roleHandler) enable(ctx *gin.Context) {
	var req dto.RoleIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	if err := h.RoleService.Enable(goctx, id); err != nil {
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
