package gtemplate

// ShowType 前端显示类型枚举
type ShowType int

const (
	ShowInput ShowType = iota
	ShowTextarea
	ShowRadio
	ShowDate
)

// -------------------------------------------------------------------------------
// -------------------------------------------------------------------------------
// -------------------------------------------------------------------------------

// QueryOp 定义查询操作的“枚举”类型
type QueryOp int

// 定义与查询操作相关的常量
const (
	Equal QueryOp = iota
	NotEqual
	GreaterThan
	GreaterOrEqual
	LessThan
	LessOrEqual
	Like
	Between
)

// queryOpLabels 将查询操作的常量映射到其对应的标签
var queryOpLabels = map[QueryOp]string{
	Equal:          "=",
	NotEqual:       "!=",
	GreaterThan:    ">",
	GreaterOrEqual: ">=",
	LessThan:       "<",
	LessOrEqual:    "<=",
	Like:           "LIKE",
	Between:        "BETWEEN",
}

// String 方法返回 QueryOp 的字符串表示，实现了 fmt.Stringer 接口
func (op QueryOp) String() string {
	return queryOpLabels[op]
}

// -------------------------------------------------------------------------------
// -------------------------------------------------------------------------------
// -------------------------------------------------------------------------------

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

// -------------------------------------------------------------------------------
// -------------------------------------------------------------------------------
// -------------------------------------------------------------------------------

type GoType int

const (
	GoInt GoType = iota
	GoUint
	GoInt64
	GoUint64
	GoFloat64
	GoStringType
	GoByte
	GoBool
	GoTime
	GoAny
)

// String 返回GoType的字符串表示
func (t GoType) String() string {
	switch t {
	case GoInt:
		return "int"
	case GoUint:
		return "uint"
	case GoInt64:
		return "int64"
	case GoUint64:
		return "uint64"
	case GoFloat64:
		return "float64"
	case GoStringType:
		return "string"
	case GoByte:
		return "byte"
	case GoBool:
		return "bool"
	case GoTime:
		return "time"
	case GoAny:
		return "any"
	default:
		return "interface{}"
	}
}
