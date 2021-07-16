package service

import (
	"context"
	"wx/app/component"
	"wx/app/dto"
	"wx/app/iface"
	"wx/app/model"
	"wx/app/zconst"

	"github.com/spf13/cast"
)

type roleService struct {
	RoleRepository       iface.RoleRepository
	PermissionRepository iface.PermissionRepository
}

func NewRoleService(r iface.RoleRepository, p iface.PermissionRepository) iface.RoleService {
	return &roleService{
		RoleRepository:       r,
		PermissionRepository: p,
	}
}

func (s *roleService) Get(ctx context.Context, id int64) (*model.Role, error) {
	return s.RoleRepository.FindByID(ctx, id)
}

func (s *roleService) List(ctx context.Context, where dto.RoleSearchReq) (dto.RoleListResp, error) {
	return s.RoleRepository.FindByWhere(ctx, where)
}

func (s *roleService) Create(ctx context.Context, m *model.Role) (*model.Role, error) {
	result, err := s.RoleRepository.Create(ctx, m)
	_, err = component.AddRoleForUser("-", m.Name)
	s.updatePermissinForRole(ctx, m.Permissions_ids, m.Name)
	return result, err
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

// 更新角色权限
func (s *roleService) updatePermissinForRole(ctx context.Context, permissionIds []int, roleAccount string) {
	for _, pid := range permissionIds {
		result, _ := s.PermissionRepository.FindByID(ctx, cast.ToInt64(pid))
		component.AddPermissionForUser(result.Url, result.Method, roleAccount)
	}
}
