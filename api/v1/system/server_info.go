package system

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type ServerInfoApi struct {
	service service.IServerInfoService
}

func NewServerInfoApi(service service.IServerInfoService) *ServerInfoApi {
	return &ServerInfoApi{
		service: service,
	}
}

func (m ServerInfoApi) GetServerInfo(_ *gin.Context, _ gdto.EmptyReq) (any, error) {
	info, err := m.service.ServerInfo()
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "获取成功",
		Data:    info,
	}, nil
}
