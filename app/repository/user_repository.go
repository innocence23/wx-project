package repository

import (
	"context"
	"log"
	"wx/app/iface"
	"wx/app/model"
	"wx/app/zconst"
	"wx/app/zerror"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) iface.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	user := &model.User{}
	if result := r.DB.Find(user, id); result.Error != nil || result.RowsAffected == 0 || user.Status == zconst.DisableStatus {
		log.Printf("数据查询失败， ID: %v. 失败原因: %v，影响行数:%d\n", id, result.Error, result.RowsAffected)
		return user, zerror.NewNotFound("uid", cast.ToString(id))
	}
	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	if result := r.DB.Where("email = ?", email).First(user); result.Error != nil || result.RowsAffected == 0 || user.Status == zconst.DisableStatus {
		log.Printf("数据查询失败， Email: %v. 失败原因: %v，影响行数:%d\n", email, result.Error, result.RowsAffected)
		return user, zerror.NewNotFound("email", email)
	}
	return user, nil
}

func (r *userRepository) Create(ctx context.Context, u *model.User) error {
	if result := r.DB.Create(u); result.Error != nil || result.RowsAffected == 0 {
		log.Printf("数据添加失败: %v. 失败原因: %v，影响行数:%d\n", u, result.Error, result.RowsAffected)
		return zerror.NewInternal()
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, u *model.User) error {
	data := model.User{
		Name:   u.Name,
		Avatar: u.Avatar,
	}
	if result := r.DB.Model(u).Updates(data); result.Error != nil || result.RowsAffected == 0 {
		log.Printf("数据更新失败: %v. 失败原因: %v，影响行数:%d\n", u, result.Error, result.RowsAffected)
		return zerror.NewInternal()
	}
	return nil
}

func (r *userRepository) UpdateStatus(ctx context.Context, id int64, status int) error {
	if result := r.DB.Model(&model.User{}).Where("id = ?", id).Update("status", status); result.Error != nil || result.RowsAffected == 0 {
		log.Printf("状态更新失败: ID:%v;Status%v. 失败原因: %v，影响行数:%d\n", id, status, result.Error, result.RowsAffected)
		return zerror.NewInternal()
	}
	return nil
}
