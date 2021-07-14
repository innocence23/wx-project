package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"wx/app/model"
	"wx/app/zerror"

	"golang.org/x/crypto/scrypt"
)

type UserService struct {
	UserRepository model.UserRepository
}

func NewUserService(ur model.UserRepository) model.UserService {
	return &UserService{
		UserRepository: ur,
	}
}

func (s *UserService) Get(ctx context.Context, id int64) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, id)
	return u, err
}

func (s *UserService) Signup(ctx context.Context, u *model.User) error {
	pw, err := hashPassword(u.Password)
	if err != nil {
		log.Printf("Unable to signup user for email: %v\n", u.Email)
		return zerror.NewInternal()
	}
	u.Password = pw
	if err := s.UserRepository.Create(ctx, u); err != nil {
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
		return false, fmt.Errorf("Unable to verify user password")
	}
	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)
	return hex.EncodeToString(shash) == pwsalt[0], nil
}
