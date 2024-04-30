package dto

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"time"
)

type TableInfoRes struct {
	TableName    string `json:"tableName"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
	TableComment string `json:"tableComment"`
	ModelName    string `json:"modelName,omitempty"`
}

type TableInfoSearchReq struct {
	gdto.PaginateReq        //分页
	TableName        string `json:"tableName" form:"tableName"`
	TableComment     string `json:"tableComment" form:"tableComment"`
	Time             []time.Time
}

type GenListRes struct {
	ID           uint      `json:"id"`
	TableNamed   string    `json:"tableNamed"`
	TableComment string    `json:"tableComment"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type GenCodeRes struct {
	GoApiCode     string `json:"goApiCode"`
	GoDtoCode     string `json:"goDtoCode"`
	GoLogicCode   string `json:"goLogicCode"`
	GoModelCode   string `json:"goModelCode"`
	GoRouterCode  string `json:"goRouterCode"`
	GoServiceCode string `json:"goServiceCode"`
	VueApiCode    string `json:"vueApiCode"`
	VueFormCode   string `json:"vueFormCode"`
	VueHookCode   string `json:"vueHookCode"`
	VueIndexCode  string `json:"vueIndexCode"`
	VueRuleCode   string `json:"vueRuleCode"`
	VueTypesCode  string `json:"vueTypesCode"`
}
