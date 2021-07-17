package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"wx/app/dto"
	"wx/app/iface"
	"wx/app/model"
	"wx/app/zconst"
	"wx/app/zerror"

	"github.com/spf13/cast"
	"golang.org/x/crypto/scrypt"
)

const password = "5nNE4zK*LLJQF&3x6D" //重置默认密码

type userService struct {
	UserRepository iface.UserRepository

	RoleRepository       iface.RoleRepository
	PermissionRepository iface.PermissionRepository
	MenuRepository       iface.MenuRepository
}

func NewUserService(r iface.UserRepository, r1 iface.RoleRepository, r2 iface.PermissionRepository, r3 iface.MenuRepository) iface.UserService {
	return &userService{
		UserRepository:       r,
		RoleRepository:       r1,
		PermissionRepository: r2,
		MenuRepository:       r3,
	}
}

func (s *userService) Find(ctx context.Context, id int64) (*model.User, error) {
	return s.UserRepository.FindByID(ctx, id)
}

func (s *userService) Signup(ctx context.Context, u *model.User) error {
	pwd, err := hashPassword(u.Password)
	if err != nil {
		log.Printf("密码加密错误 email: %v\n", u.Email)
		return zerror.NewInternal()
	}
	u.Password = pwd
	if err := s.UserRepository.Create(ctx, u); err != nil {
		return err
	}
	return nil
}

func (s *userService) Signin(ctx context.Context, u *model.User) error {
	user, err := s.UserRepository.FindByEmail(ctx, u.Email)
	if err != nil {
		log.Printf("用户邮箱不存在 email: %v\n", u.Email)
		return err
	}
	match, err := comparePasswords(user.Password, u.Password)
	if err != nil {
		return zerror.NewInternal()
	}
	if !match {
		return zerror.NewAuthorization("邮箱或密码错误")
	}
	*u = *user //此处可直接取用户详情
	return nil
}

func (s *userService) List(ctx context.Context, where dto.UserSearchReq) (dto.UserListResp, error) {
	return s.UserRepository.FindByWhere(ctx, where)
}

func (s *userService) Update(ctx context.Context, u *model.User) error {
	return s.UserRepository.Update(ctx, u)
}

func (s *userService) Disable(ctx context.Context, id int64) error {
	return s.UserRepository.UpdateStatus(ctx, id, zconst.DisableStatus)
}

func (s *userService) Enable(ctx context.Context, id int64) error {
	return s.UserRepository.UpdateStatus(ctx, id, zconst.NormalStatus)
}

func (s *userService) Resetpwd(ctx context.Context, id int64) error {
	pwd, err := hashPassword(password)
	if err != nil {
		log.Printf("密码加密错误 email: %v\n", id)
		return zerror.NewInternal()
	}
	u := &model.User{
		ID:       id,
		Password: pwd,
	}
	u.Password = pwd
	if err := s.UserRepository.UpdatePassword(ctx, u); err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))
	return hashedPW, nil
}

func comparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	pwsalt := strings.Split(storedPassword, ".")
	salt, err := hex.DecodeString(pwsalt[1])
	if err != nil {
		return false, fmt.Errorf("无效密码")
	}
	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)
	return hex.EncodeToString(shash) == pwsalt[0], nil
}

//--- 权限
func (s *userService) GetMenus(ctx context.Context, roleIds []int64) []model.Menu {
	var roleList []model.Role
	var menuList []model.Menu
	var menu_ids []int64
	var level int64 = 0
	roleList, _ = s.RoleRepository.FindByIds(ctx, roleIds)
	for _, role := range roleList {
		menu_ids = append(menu_ids, role.MenuIds...)
	}
	menuList, _ = s.MenuRepository.FindByIds(ctx, menu_ids)
	for _, item := range menuList {
		if item.PId > level {
			menu_ids = append(menu_ids, item.PId)
		}
	}
	menuList, _ = s.MenuRepository.FindByIds(ctx, menu_ids)
	menuList = GetMenuTreeRouter(menuList, level)
	return menuList
}

func GetMenuTreeRouter(menus []model.Menu, pid int64) []model.Menu {
	var list []model.Menu
	for _, v := range menus {
		if v.PId == pid {
			v.Children = GetMenuTreeRouter(menus, cast.ToInt64(v.Id))
			list = append(list, v)
		}
	}
	return list
}
