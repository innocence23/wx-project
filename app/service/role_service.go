package service

import (
	"context"
	"wx/app/dto"
	"wx/app/iface"
	"wx/app/model"
	"wx/app/zconst"
)

type roleService struct {
	RoleRepository iface.RoleRepository
}

func NewRoleService(r iface.RoleRepository) iface.RoleService {
	return &roleService{
		RoleRepository: r,
	}
}

func (s *roleService) Get(ctx context.Context, id int64) (*model.Role, error) {
	return s.RoleRepository.FindByID(ctx, id)
}

func (s *roleService) List(ctx context.Context, where dto.RoleSearchReq) (dto.RoleListResp, error) {
	return s.RoleRepository.FindByWhere(ctx, where)
}

func (s *roleService) Create(ctx context.Context, m *model.Role) (*model.Role, error) {
	return s.RoleRepository.Create(ctx, m)
}

func (s *roleService) Update(ctx context.Context, m *model.Role) error {
	return s.RoleRepository.Update(ctx, m)
}

func (s *roleService) Disable(ctx context.Context, id int64) error {
	return s.RoleRepository.UpdateStatus(ctx, id, zconst.DisableStatus)
}

func (s *roleService) Enable(ctx context.Context, id int64) error {
	return s.RoleRepository.UpdateStatus(ctx, id, zconst.NormalStatus)
}
