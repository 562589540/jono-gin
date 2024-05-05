package enum

type TsType int

const (
	TsStringType TsType = iota
	TsNumberType
	TsAnyType
)

// String 返回TsType的字符串表示
func (t TsType) String() string {
	switch t {
	case TsStringType:
		return "string"
	case TsNumberType:
		return "number"
	case TsAnyType:
		return "any"
	default:
		return "unknown"
	}
}
