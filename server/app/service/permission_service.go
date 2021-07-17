package service

import (
	"context"
	"wx/app/dto"
	"wx/app/iface"
	"wx/app/model"
	"wx/app/zconst"
)

type permissionService struct {
	PermissionRepository iface.PermissionRepository
}

func NewPermissionService(r iface.PermissionRepository) iface.PermissionService {
	return &permissionService{
		PermissionRepository: r,
	}
}

func (s *permissionService) Find(ctx context.Context, id int64) (*model.Permission, error) {
	return s.PermissionRepository.FindByID(ctx, id)
}

func (s *permissionService) List(ctx context.Context, where dto.PermissionSearchReq) (dto.PermissionListResp, error) {
	return s.PermissionRepository.FindByWhere(ctx, where)
}

func (s *permissionService) Create(ctx context.Context, m *model.Permission) (*model.Permission, error) {
	return s.PermissionRepository.Create(ctx, m)
}

func (s *permissionService) Update(ctx context.Context, m *model.Permission) error {
	return s.PermissionRepository.Update(ctx, m)
}

func (s *permissionService) Disable(ctx context.Context, id int64) error {
	return s.PermissionRepository.UpdateStatus(ctx, id, zconst.DisableStatus)
}

func (s *permissionService) Enable(ctx context.Context, id int64) error {
	return s.PermissionRepository.UpdateStatus(ctx, id, zconst.NormalStatus)
}

func (s *permissionService) AutoGenerate(ctx context.Context, routers []map[string]string) error {
	for _, router := range routers {
		res, _ := s.PermissionRepository.FindByUrlAndMethod(ctx, router["url"], router["method"])
		if res.Id == 0 {
			m := &model.Permission{
				Name:   router["name"],
				Group:  router["name"],
				Url:    router["url"],
				Method: router["method"],
			}
			_, err := s.PermissionRepository.Create(ctx, m)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
