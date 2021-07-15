package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"wx/app/model"
	"wx/app/zconst"
	"wx/app/zerror"

	"golang.org/x/crypto/scrypt"
)

type userService struct {
	UserRepository model.UserRepository
}

func NewUserService(ur model.UserRepository) model.UserService {
	return &userService{
		UserRepository: ur,
	}
}

func (s *userService) Get(ctx context.Context, id int64) (*model.User, error) {
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

func (s *userService) UpdateDetail(ctx context.Context, u *model.User) error {
	return s.UserRepository.Update(ctx, u)
}

func (s *userService) DisableUser(ctx context.Context, id int64) error {
	return s.UserRepository.UpdateStatus(ctx, id, zconst.DisableStatus)
}

func (s *userService) EnableUser(ctx context.Context, id int64) error {
	return s.UserRepository.UpdateStatus(ctx, id, zconst.NormalStatus)
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
