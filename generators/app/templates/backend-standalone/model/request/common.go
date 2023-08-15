package request

type SingleIdStruct struct {
	ID uint `form:"id" binding:"required"`
}

type SingleJsonStruct struct {
	ID uint `json:"id" binding:"required"`
}

type UriIdStruct struct {
	ID uint `uri:"id" binding:"required"`
}

type UriPointerId struct {
	ID *uint `uri:"id" binding:"required"`
}

type ListParamsStruct struct {
	Page  int `form:"page" binding:"required"`
	Limit int `form:"limit" binding:"required"`
}

type ListJsonStruct struct {
	Page  int `json:"page" binding:"required"`
	Limit int `json:"limit" binding:"required"`
}
