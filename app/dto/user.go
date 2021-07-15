package dto

type SignupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

type SigninReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

type DetailsReq struct {
	Name   string `json:"name" binding:"omitempty,max=50"`
	Avatar string `json:"avatar" binding:"required"`
}