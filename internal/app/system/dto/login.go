package dto

import "github.com/mileusna/useragent"

type AdminLoginReq struct {
	UserName string `json:"username,omitempty" binding:"required" m:"账号不能为空"`
	Password string `json:"password" binding:"password" m:"密码格式应为8-18位数字、字母、符号的任意两种组合"`
}

type AdminLoginRes struct {
	UserName     string   `json:"username"`
	NickName     string   `json:"nickname"`
	Avatar       string   `json:"avatar"`
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	Expires      int64    `json:"expires"`
	Roles        []string `json:"roles"`
}

type RefreshTokenRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Expires      int64  `json:"expires"`
}

type LoginParam struct {
	Ip       string
	UserName string
	Behavior string
	Status   bool
	UserId   uint
	Ua       useragent.UserAgent
}
