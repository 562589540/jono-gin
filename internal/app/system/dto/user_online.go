package dto

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"time"
)

type UserOnlineSearchReq struct {
	gdto.PaginateReq        //分页
	UserName         string `json:"username,omitempty" form:"username"`
}

type UserOnline struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"username"`
	Address   string    `json:"address"`
	Browser   string    `json:"browser"`
	Ip        string    `json:"ip"`
	System    string    `json:"system"`
	LoginTime time.Time `json:"loginTime"`
}
