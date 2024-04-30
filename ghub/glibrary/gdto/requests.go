package gdto

type EmptyReq struct{}

type IDReq struct {
	ID uint `json:"id" form:"id" uri:"id"  binding:"required" m:"查询ID不能为空"`
}

type IDSReq struct {
	IDS []uint `json:"ids" form:"ids" uri:"ids"  binding:"required" m:"查询ID不能为空"`
}

type PaginateReq struct {
	Page  int `json:"page,omitempty"  form:"page"`
	Limit int `json:"limit,omitempty"  form:"limit"`
}

func (m *PaginateReq) GetPage() int {
	if m.Page <= 0 {
		m.Page = 1
	}
	return m.Page
}

func (m *PaginateReq) GetOffset() int {
	return (m.GetPage() - 1) * m.GetLimit()
}

func (m *PaginateReq) GetLimit() int {
	if m.Limit <= 0 {
		m.Limit = 10
	}
	return m.Limit
}
