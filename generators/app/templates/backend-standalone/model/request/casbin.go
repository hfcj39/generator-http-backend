package request

type UpdateCasbin struct {
	OldPolicy []string `json:"oldPolicy" binding:"required"`
	NewPolicy []string `json:"newPolicy" binding:"required"`
}

type UpdateCasbinById struct {
	ID uint   `json:"id" binding:"required"`
	V0 string `json:"V0" binding:"required"`
	V2 string `json:"V2" binding:"required"`
	V3 string `json:"V3" binding:"required"`
}
