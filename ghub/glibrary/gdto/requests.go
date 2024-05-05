package gdto

import "strconv"

type EmptyReq struct{}

type CaptchaRes struct {
	ID   string `json:"id"`
	B64s string `json:"b64S"`
}

type CheckCaptchaRes struct {
	Success bool `json:"success"`
}

type CheckCaptchaReq struct {
	ID      string `json:"id"`
	Captcha string `json:"captcha"`
}

type IDReq struct {
	ID uint `json:"id" form:"id" uri:"id"  binding:"required" m:"查询ID不能为空"`
}

type IDSReq struct {
	IDS []uint `json:"ids" form:"ids" uri:"ids"  binding:"required" m:"查询ID不能为空"`
}

type SnowflakeReq struct {
	ID string `json:"id" form:"id" uri:"id"  binding:"required" m:"查询ID不能为空"`
}

func (s *SnowflakeReq) Int64ID() int64 {
	idInt64, err := strconv.ParseInt(s.ID, 10, 64)
	if err != nil {
		return 0
	}
	return idInt64
}

type SnowflakeSReq struct {
	IDS []string `json:"ids" form:"ids" uri:"ids"  binding:"required" m:"查询ID不能为空"`
}

func (s *SnowflakeSReq) Int64IDS() []int64 {
	var idsInt64 []int64
	for _, idStr := range s.IDS {
		idInt64, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			continue
		}
		idsInt64 = append(idsInt64, idInt64)
	}
	return idsInt64
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
