package dto

type RoleIdReq struct {
	ID int64 `json:"id" form:"id" binding:"required" label:"角色ID"`
}

type RoleCreateReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	MenuIds     string `json:"menu_ids" binding:"required"`
}

type RoleUpdateReq struct {
	ID          int64  `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	MenuIds     string `json:"menu_ids" binding:"required"`
}

type RoleSearchReq struct {
	Name         string `json:"name"`
	CreatedAtMin string `json:"created_at_min"`
	CreatedAtMax string `json:"created_at_max"`
	Pagination
}
