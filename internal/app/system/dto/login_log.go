package dto

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"time"
)

type LoginLogSearchReq struct {
	gdto.PaginateReq
	UserName  string `json:"username" form:"username"`
	Status    string `json:"status" form:"status"`
	LoginTime []time.Time
}

type LoginLogAddReq struct {
}

type LoginLog struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"username"`
	Address   string    `json:"address"`
	Behavior  string    `json:"behavior"`
	Browser   string    `json:"browser"`
	Ip        string    `json:"ip"`
	Status    int       `json:"status"`
	System    string    `json:"system"`
	LoginTime time.Time `json:"loginTime"`
}
