package handler

import (
	"log"
	"wx/app/dto"
	"wx/app/handler/commonhandler"
	"wx/app/iface"
	"wx/app/model"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type menuHandler struct {
	MenuService iface.MenuService
}

func NewMenuHandler(s iface.MenuService) *menuHandler {
	return &menuHandler{
		MenuService: s,
	}
}

func (h *menuHandler) Router(router *gin.RouterGroup) {
	grouter := router.Group("/menu")
	grouter.GET("/info", h.show)
	grouter.POST("/list", h.list) //参数多，改为post方便
	grouter.POST("", h.create)
	grouter.PUT("", h.update)
	grouter.PUT("/enable", h.enable)
	grouter.PUT("/disable", h.disable)
}

func (h *menuHandler) show(ctx *gin.Context) {
	var req dto.MenuIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	result, err := h.MenuService.Find(goctx, id)
	if err != nil {
		log.Printf("信息不存在: %v \n%v", id, err)
		e := zerror.NewNotFound("menu", cast.ToString(id))
		commonhandler.Fail(ctx, e.Status(), gin.H{
			"error": e,
		})
		return
	}
	commonhandler.Success(ctx, result)
}

func (h *menuHandler) list(ctx *gin.Context) {
	var req dto.MenuSearchReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	where := req
	goctx := ctx.Request.Context()
	list, err := h.MenuService.List(goctx, where)
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

func (h *menuHandler) create(ctx *gin.Context) {
	var req dto.MenuCreateReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	data := &model.Menu{
		PId:         *req.PId,
		Name:        req.Name,
		Description: req.Description,
		Url:         req.Url,
		Icon:        req.Icon,
		Weight:      req.Weight,
	}
	goctx := ctx.Request.Context()
	result, err := h.MenuService.Create(goctx, data)
	if err != nil {
		log.Printf("数据创建失败: %v\n", err.Error())
		commonhandler.Fail(ctx, zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}
	commonhandler.Success(ctx, result)
}

func (h *menuHandler) update(ctx *gin.Context) {
	var req dto.MenuUpdateReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	data := &model.Menu{
		Id:          req.ID,
		PId:         *req.PId,
		Name:        req.Name,
		Description: req.Description,
		Url:         req.Url,
		Icon:        req.Icon,
		Weight:      req.Weight,
	}
	goctx := ctx.Request.Context()
	err := h.MenuService.Update(goctx, data)
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

func (h *menuHandler) disable(ctx *gin.Context) {
	var req dto.MenuIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	if err := h.MenuService.Disable(goctx, id); err != nil {
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

func (h *menuHandler) enable(ctx *gin.Context) {
	var req dto.MenuIdReq
	if ok := commonhandler.BindData(ctx, &req); !ok {
		return
	}
	id := req.ID
	goctx := ctx.Request.Context()
	if err := h.MenuService.Enable(goctx, id); err != nil {
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
