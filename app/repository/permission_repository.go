package repository

import (
	"context"
	"log"
	"wx/app/dto"
	"wx/app/iface"
	"wx/app/model"
	"wx/app/zconst"
	"wx/app/zerror"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type permissionRepository struct {
	DB *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) iface.PermissionRepository {
	return &permissionRepository{
		DB: db,
	}
}

func (r *permissionRepository) FindByID(ctx context.Context, id int64) (*model.Permission, error) {
	permission := &model.Permission{}
	if result := r.DB.Find(permission, id); result.Error != nil || result.RowsAffected == 0 || permission.Status == zconst.DisableStatus {
		log.Printf("数据查询失败， ID: %v. 失败原因: %v，影响行数:%d\n", id, result.Error, result.RowsAffected)
		return permission, zerror.NewNotFound("uid", cast.ToString(id))
	}
	return permission, nil
}

func (r *permissionRepository) FindByWhere(ctx context.Context, where dto.PermissionSearchReq) (dto.PermissionListResp, error) {
	permissions := make([]model.Permission, 0)
	query := r.DB.Model(&model.Permission{})
	if len(where.Name) > 0 {
		query = query.Where("name like ?", "%"+where.Name+"%")
	}
	if len(where.Url) > 0 {
		query = query.Where("url like ?", "%"+where.Url+"%")
	}
	if len(where.CreatedAtMin) > 0 {
		query = query.Where("created_at >= ?", where.CreatedAtMin)
		query = query.Where("created_at <= ?", where.CreatedAtMax)
	}
	var total int64 = 0
	if err := query.Count(&total).Error; err != nil {
		return dto.PermissionListResp{}, zerror.NewInternal()
	}

	if err := query.Order("id DESC").Limit(where.PageSize).Offset((where.Page - 1) * where.PageSize).Find(&permissions).Error; err != nil {
		log.Printf("数据查询失败， where: %#v. 失败原因: %v\n", where, err)
		return dto.PermissionListResp{}, zerror.NewInternal()
	}
	result := dto.PermissionListResp{
		List:  permissions,
		Total: total,
		Pagination: dto.Pagination{
			Page:     where.Page,
			PageSize: where.PageSize,
		},
	}
	return result, nil
}

func (r *permissionRepository) Create(ctx context.Context, m *model.Permission) (*model.Permission, error) {
	if result := r.DB.Create(m); result.Error != nil || result.RowsAffected == 0 {
		log.Printf("数据添加失败: %v. 失败原因: %v，影响行数:%d\n", m, result.Error, result.RowsAffected)
		return nil, zerror.NewInternal()
	}
	return m, nil
}

func (r *permissionRepository) Update(ctx context.Context, m *model.Permission) error {
	data := model.Permission{
		PId:    m.PId,
		Name:   m.Name,
		Url:    m.Url,
		Method: m.Method,
	}
	result := r.DB.Model(m).Updates(data)
	if result.Error != nil {
		log.Printf("数据更新失败: %v. 失败原因: %v，影响行数:%d\n", m, result.Error, result.RowsAffected)
		return zerror.NewInternal()
	}
	if result.RowsAffected == 0 {
		return zerror.NewNotFound("id", cast.ToString(data.Id))
	}
	return nil
}

func (r *permissionRepository) UpdateStatus(ctx context.Context, id int64, status int) error {
	m := &model.Permission{Id: id}
	result := r.DB.Model(m).Update("status", status)
	if result.Error != nil {
		log.Printf("状态更新失败: ID:%v;Status%v. 失败原因: %v，影响行数:%d\n", id, status, result.Error, result.RowsAffected)
		return zerror.NewInternal()
	}
	if result.RowsAffected == 0 {
		return zerror.NewNotFound("id", cast.ToString(id))
	}
	return nil
}
