package dto

type RoleIdReq struct {
	ID int64 `json:"id" binding:"required"`
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
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Pagination
}
