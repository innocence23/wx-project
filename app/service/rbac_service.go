package service

import (
	"wx/app/iface"
)

type rbacService struct {
}

func NewRbacService(r iface.UserRepository) *rbacService {
	return &rbacService{}
}
