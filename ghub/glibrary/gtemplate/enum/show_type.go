package enum

// ShowType 前端显示类型枚举
type ShowType int

const (
	ShowInput ShowType = iota
	ShowTextarea
	ShowRadio
	ShowDate
	ShowNum
)
