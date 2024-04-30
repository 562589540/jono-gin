package system

import (
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type LoginLogApi struct {
	service service.ILoginLogService
}

func NewLoginLogApi(service service.ILoginLogService) *LoginLogApi {
	return &LoginLogApi{
		service: service,
	}
}
func (m LoginLogApi) Delete(c *gin.Context, req gdto.IDSReq) (any, error) {
	if err := m.service.Delete(c, req.IDS); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}

func (m LoginLogApi) DeleteAll(c *gin.Context, _ gdto.EmptyReq) (any, error) {
	if err := m.service.DeleteAll(c); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}

func (m LoginLogApi) List(c *gin.Context, req dto.LoginLogSearchReq) (any, error) {
	req.LoginTime, _ = ghub.ParseTimeInterval(c, "loginTime")
	list, total, err := m.service.List(c, &req)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "查询成功",
		Data: gdto.ListRes{
			List:  list,
			Total: total,
		},
	}, nil
}
