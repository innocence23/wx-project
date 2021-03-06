package service

import (
	"context"
	"wx/app/dto"
	"wx/app/iface"
	"wx/app/model"
	"wx/app/zconst"
)

type menuService struct {
	MenuRepository       iface.MenuRepository
	RoleRepository       iface.RoleRepository
	PermissionRepository iface.PermissionRepository
}

func NewMenuService(r iface.MenuRepository, r1 iface.RoleRepository, r2 iface.PermissionRepository) iface.MenuService {
	return &menuService{
		MenuRepository:       r,
		RoleRepository:       r1,
		PermissionRepository: r2,
	}
}

func (s *menuService) Find(ctx context.Context, id int64) (*model.Menu, error) {
	return s.MenuRepository.FindByID(ctx, id)
}

func (s *menuService) List(ctx context.Context, where dto.MenuSearchReq) (dto.MenuListResp, error) {
	return s.MenuRepository.FindByWhere(ctx, where)
}

func (s *menuService) Create(ctx context.Context, m *model.Menu) (*model.Menu, error) {
	return s.MenuRepository.Create(ctx, m)
}

func (s *menuService) Update(ctx context.Context, m *model.Menu) error {
	return s.MenuRepository.Update(ctx, m)
}

func (s *menuService) Disable(ctx context.Context, id int64) error {
	return s.MenuRepository.UpdateStatus(ctx, id, zconst.DisableStatus)
}

func (s *menuService) Enable(ctx context.Context, id int64) error {
	return s.MenuRepository.UpdateStatus(ctx, id, zconst.NormalStatus)
}
