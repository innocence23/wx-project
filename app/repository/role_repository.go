package repository

import (
	"context"
	"log"
	"wx/app/dto"
	"wx/app/model"
	"wx/app/zconst"
	"wx/app/zerror"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type roleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) model.RoleRepository {
	return &roleRepository{
		DB: db,
	}
}

func (r *roleRepository) FindByID(ctx context.Context, id int64) (*model.Role, error) {
	role := &model.Role{}
	if result := r.DB.Find(role, id); result.Error != nil || result.RowsAffected == 0 || role.Status == zconst.DisableStatus {
		log.Printf("数据查询失败， ID: %v. 失败原因: %v，影响行数:%d\n", id, result.Error, result.RowsAffected)
		return role, zerror.NewNotFound("uid", cast.ToString(id))
	}
	return role, nil
}

func (r *roleRepository) FindByWhere(ctx context.Context, where dto.RoleSearchReq) ([]model.Role, error) {
	roles := make([]model.Role, 0)
	x := r.DB
	if len(where.Name) > 0 {
		x = x.Where("name like ?", "%"+where.Name+"%")
	}
	if len(where.CreatedAtMin) > 0 {
		x = x.Where("created_at >= ?", where.CreatedAtMin)
		x = x.Where("created_at <= ?", where.CreatedAtMax)
	}

	if err := x.Order("id DESC").Limit(where.PageSize).Offset((where.Page - 1) * where.PageSize).Find(&roles).Error; err != nil {
		log.Printf("数据查询失败， where: %#v. 失败原因: %v\n", where, err)
		return nil, zerror.NewInternal()
	}
	return roles, nil
}

func (r *roleRepository) Create(ctx context.Context, m *model.Role) (*model.Role, error) {
	if result := r.DB.Create(m); result.Error != nil || result.RowsAffected == 0 {
		log.Printf("数据添加失败: %v. 失败原因: %v，影响行数:%d\n", m, result.Error, result.RowsAffected)
		return nil, zerror.NewInternal()
	}
	return m, nil
}

func (r *roleRepository) Update(ctx context.Context, m *model.Role) error {
	data := model.Role{
		Name:        m.Name,
		Description: m.Description,
		MenuIds:     m.MenuIds,
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

func (r *roleRepository) UpdateStatus(ctx context.Context, id int64, status int) error {
	m := &model.Role{Id: id}
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
