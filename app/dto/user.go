package dto

type SignupReq struct {
	Email    string `json:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" binding:"required,gte=6,lte=30" label:"密码"`
}

type SigninReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

type DetailsReq struct {
	Name   string `json:"name" binding:"omitempty,max=50"`
	Avatar string `json:"avatar" binding:"required"`
}

type UserJWT struct {
	ID     int64
	Name   string
	Email  string
	Avatar string
}
