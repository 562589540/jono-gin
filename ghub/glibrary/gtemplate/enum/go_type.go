package enum

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
