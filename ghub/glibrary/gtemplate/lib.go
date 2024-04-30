package gtemplate

import "strings"

// ConvertDBTypeToTS 将数据库字段类型转换为TypeScript类型
func ConvertDBTypeToTS(mysqlType string) TsType {
	// 根据类型前缀映射到TypeScript类型
	switch {
	case strings.HasPrefix(mysqlType, "bigint"):
		return TsNumberType
	case strings.HasPrefix(mysqlType, "int"),
		strings.HasPrefix(mysqlType, "smallint"),
		strings.HasPrefix(mysqlType, "tinyint"),
		strings.HasPrefix(mysqlType, "mediumint"):
		// 特殊处理 tinyint(1) 通常用于表示布尔值
		if mysqlType == "tinyint(1)" {
			return TsNumberType
		}
		return TsNumberType
	case strings.HasPrefix(mysqlType, "float"),
		strings.HasPrefix(mysqlType, "double"),
		strings.HasPrefix(mysqlType, "decimal"):
		return TsNumberType
	case strings.HasPrefix(mysqlType, "char"),
		strings.HasPrefix(mysqlType, "varchar"),
		strings.HasPrefix(mysqlType, "text"),
		strings.HasPrefix(mysqlType, "longtext"):
		return TsStringType
	case strings.HasPrefix(mysqlType, "enum"),
		strings.HasPrefix(mysqlType, "set"):
		return TsStringType // 或者可以用更具体的联合类型，需要额外的逻辑来处理
	case strings.HasPrefix(mysqlType, "date"),
		strings.HasPrefix(mysqlType, "datetime"),
		strings.HasPrefix(mysqlType, "timestamp"),
		strings.HasPrefix(mysqlType, "time"),
		strings.HasPrefix(mysqlType, "year"):
		return TsStringType // 或者 "string" 如果你想用ISO字符串
	case strings.HasPrefix(mysqlType, "bit"):
		return TsNumberType
	default:
		// 未知或者较少用到的类型，默认返回 "any"
		return TsAnyType
	}
}

// ConvertDBTypeToGoType 将数据库类型转换为 Go 类型
// 该函数目前支持常见的 MySQL 数据类型
func ConvertDBTypeToGoType(dbType string) GoType {
	// 将数据库类型统一处理为小写，以便比较
	t := strings.ToLower(dbType)

	// 根据不同的数据库类型返回相应的 Go 类型
	switch {
	case strings.Contains(t, "bigint"):
		return GoUint
	case strings.Contains(t, "int"):
		return GoInt
	case strings.Contains(t, "varchar"), strings.Contains(t, "text"), strings.Contains(t, "char"):
		return GoStringType
	case strings.Contains(t, "bool"):
		return GoBool
	case strings.Contains(t, "float"), strings.Contains(t, "double"):
		return GoFloat64
	case strings.Contains(t, "decimal"):
		return GoFloat64 // 或使用更精确的 decimal 包
	case strings.Contains(t, "date"), strings.Contains(t, "time"):
		return GoTime // 需要 import "time"
	default:
		return GoAny // 未知类型
	}
}
