package enum

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
