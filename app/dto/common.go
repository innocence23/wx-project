package dto

type Pagination struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"page_size" binding:"required"`
}
