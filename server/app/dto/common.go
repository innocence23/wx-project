package dto

type Pagination struct {
	Page     int `json:"page" binding:"required" label:"当前页"`
	PageSize int `json:"page_size" binding:"required" label:"分页个数"`
}
