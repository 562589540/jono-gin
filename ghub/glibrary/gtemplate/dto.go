package gtemplate

type TableFields struct {
	Field     string   `json:"field"`     // 数据库字段名
	FieldDes  string   `json:"fieldDes"`  // 字段描述
	MysqlType string   `json:"mysqlType"` // MySQL字段类型
	GoType    GoType   `json:"goType"`    // Go语言中的类型
	TsType    TsType   `json:"tsType"`    // TypeScript中的类型
	GoName    string   `json:"goName"`    // Go中的字段名（大驼峰）
	JsonName  string   `json:"jsonName"`  // JSON键名（小驼峰或蛇形）
	Edit      bool     `json:"edit"`      // 是否用于编辑
	List      bool     `json:"list"`      // 是否在列表中显示
	FillUp    bool     `json:"fillUp"`    // 是否占一行
	Details   bool     `json:"details"`   // 是否在详情页显示
	Query     bool     `json:"query"`     // 是否可查询
	QueryType QueryOp  `json:"queryType"` // 查询类型，例如 "equals", "range", etc.
	Required  bool     `json:"required"`  // 是否必填
	ShowType  ShowType `json:"showType"`  // 显示类型，例如 "input", "textarea", etc.
	Date      string   `json:"date"`      // 可用于特定日期处理
}

type BaseInfo struct {
	TableName    string   `json:"tableName" binding:"required"`    //表名
	TableComment string   `json:"tableComment" binding:"required"` //表描述
	Author       string   `json:"author" binding:"required"`       //作者
	Config       []string `json:"config"`                          //配置
	Covers       []string `json:"covers"`                          //覆盖
	ModelName    string   `json:"modelName" binding:"required"`    //模型名称
	Remark       string   `json:"remark"`                          //备注
	SortField    string   `json:"sortField" binding:"required"`    //排序字段
	SortWay      string   `json:"sortWay" binding:"required"`      //排序类型
}

type GenInfo struct {
	Template     string `json:"template" binding:"required"`     //模版类型
	Directory    string `json:"directory" binding:"required"`    //所属目录
	BusinessName string `json:"businessName" binding:"required"` //生成业务名
	FunctionName string `json:"functionName" binding:"required"` //生成功能名 中文
	PackPath     string `json:"packPath" binding:"required"`     //生成包路径
}

type GenCodes struct {
	GoApiCode     string
	GoDtoCode     string
	GoLogicCode   string
	GoModelCode   string
	GoRouterCode  string
	GoServiceCode string
	VueApiCode    string
	VueFormCode   string
	VueHookCode   string
	VueIndexCode  string
	VueRuleCode   string
	VueTypesCode  string
}
