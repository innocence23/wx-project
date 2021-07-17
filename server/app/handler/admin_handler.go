package handler

import (
	"log"
	"wx/app/component"
	"wx/app/dto"
	"wx/app/handler/commonhandler"
	"wx/app/iface"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type adminHandler struct {
	UserService iface.UserService
}

func NewAdminHandler(s iface.UserService) *adminHandler {
	return &adminHandler{
		UserService: s,
	}
}

func (h *adminHandler) Router(router *gin.RouterGroup) {
	router = router.Group("/user")
	router.GET("/info", h.info)
	router.PUT("/enable", h.enable)
	router.PUT("/disable", h.disable)
	router.PUT("/resetpwd", h.resetpwd)
	router.POST("/list", h.list)
}

func (h *adminHandler) info(ctx *gin.Context) {
	var req dto.UserIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	result, err := h.UserService.Find(goctx, id)
	if err != nil {
		log.Printf("无法找到用户: %v %#v", id, err)
		e := zerror.NewNotFound("admin", cast.ToString(id))
		commonhandler.Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}
	commonhandler.Success(ctx, result)
}

func (h *adminHandler) disable(ctx *gin.Context) {
	var req dto.UserIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	err := h.UserService.Disable(goctx, id)
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

func (h *adminHandler) enable(ctx *gin.Context) {
	var req dto.UserIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	err := h.UserService.Enable(goctx, id)
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

func (h *adminHandler) resetpwd(ctx *gin.Context) {
	var req dto.UserIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	err := h.UserService.Resetpwd(goctx, id)
	if err != nil {
		log.Printf("用户重置密码失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"message": "操作成功",
	})
}

func (h *adminHandler) list(ctx *gin.Context) {
	var req dto.UserSearchReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	where := req
	goctx := ctx.Request.Context()
	list, err := h.UserService.List(goctx, where)
	if err != nil {
		log.Printf("无法找到数据: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, list)
}

func (h *adminHandler) setRoles(ctx *gin.Context) {
	var req dto.UserAddRolesReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	uid := cast.ToString(id)
	_, err := component.AddRolesForUser(uid, req.RoleIds)
	if err != nil {
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"message": "操作成功",
	})
}

func (h *adminHandler) getRoles(ctx *gin.Context) {
	var req dto.UserIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	uid := cast.ToString(id)
	roleIds, err := component.GetRolesForUser(uid)
	if err != nil {
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, gin.H{
		"roleids": roleIds,
	})
}
