package dto

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"time"
)

type OperLogSearchReq struct {
	gdto.PaginateReq        //分页
	Module           string `form:"module" json:"module"`
	Status           string `form:"status" json:"status"`
	CreatedAt        []time.Time
}

type OperLogUpdateReq struct {
	ID uint `json:"id" form:"id"`
}

type OperLogAddReq struct {
}

type OperLog struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"username"`
	Address   string    `json:"address"`
	Browser   string    `json:"browser"`
	Ip        string    `json:"ip"`
	Module    string    `json:"module"`
	Summary   string    `json:"summary"`
	System    string    `json:"system"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"operatingTime"`
}
