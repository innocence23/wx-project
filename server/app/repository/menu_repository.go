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

type menuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository(db *gorm.DB) iface.MenuRepository {
	return &menuRepository{
		DB: db,
	}
}

func (r *menuRepository) FindByID(ctx context.Context, id int64) (*model.Menu, error) {
	menu := &model.Menu{}
	if result := r.DB.Find(menu, id); result.Error != nil || result.RowsAffected == 0 || menu.Status == zconst.DisableStatus {
		log.Printf("数据查询失败， ID: %v. 失败原因: %v，影响行数:%d\n", id, result.Error, result.RowsAffected)
		return menu, zerror.NewNotFound("uid", cast.ToString(id))
	}
	return menu, nil
}

func (r *menuRepository) FindByIds(ctx context.Context, ids []int64) ([]model.Menu, error) {
	menu := []model.Menu{}
	if result := r.DB.Where(ids).Find(&menu); result.Error != nil || result.RowsAffected == 0 {
		log.Printf("数据查询失败， ID: %v. 失败原因: %v，影响行数:%d\n", ids, result.Error, result.RowsAffected)
		return menu, zerror.NewNotFound("uid", cast.ToString(ids))
	}
	return menu, nil
}

func (r *menuRepository) FindByWhere(ctx context.Context, where dto.MenuSearchReq) (dto.MenuListResp, error) {
	menus := make([]model.Menu, 0)
	query := r.DB.Model(&model.Menu{})
	if len(where.Name) > 0 {
		query = query.Where("name like ?", "%"+where.Name+"%")
	}
	if len(where.CreatedAtMin) > 0 {
		query = query.Where("created_at >= ?", where.CreatedAtMin)
		query = query.Where("created_at <= ?", where.CreatedAtMax)
	}
	var total int64 = 0
	if err := query.Count(&total).Error; err != nil {
		return dto.MenuListResp{}, zerror.NewInternal()
	}

	if err := query.Order("id DESC").Limit(where.PageSize).Offset((where.Page - 1) * where.PageSize).Find(&menus).Error; err != nil {
		log.Printf("数据查询失败， where: %#v. 失败原因: %v\n", where, err)
		return dto.MenuListResp{}, zerror.NewInternal()
	}
	result := dto.MenuListResp{
		List:  menus,
		Total: total,
		Pagination: dto.Pagination{
			Page:     where.Page,
			PageSize: where.PageSize,
		},
	}
	return result, nil
}

func (r *menuRepository) Create(ctx context.Context, m *model.Menu) (*model.Menu, error) {
	if result := r.DB.Create(m); result.Error != nil || result.RowsAffected == 0 {
		log.Printf("数据添加失败: %v. 失败原因: %v，影响行数:%d\n", m, result.Error, result.RowsAffected)
		return nil, zerror.NewInternal()
	}
	return m, nil
}

func (r *menuRepository) Update(ctx context.Context, m *model.Menu) error {
	data := model.Menu{
		PId:         m.PId,
		Name:        m.Name,
		Description: m.Description,
		Url:         m.Url,
		Icon:        m.Icon,
		Weight:      m.Weight,
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

func (r *menuRepository) UpdateStatus(ctx context.Context, id int64, status int) error {
	m := &model.Menu{Id: id}
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
