package service

import "github.com/562589540/jono-gin/internal/app/system/dto"

type IServerInfoService interface {
	ServerInfo() (*dto.ServerInfoRes, error)
}
