package system

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type OperLogApi struct {
	operLogService service.IOperLogService
}

func NewOperLogApi(operLogService service.IOperLogService) *OperLogApi {
	return &OperLogApi{
		operLogService: operLogService,
	}
}

func (m OperLogApi) List(c *gin.Context, req dto.OperLogSearchReq) (any, error) {
	req.CreatedAt, _ = gutils.ParseTimeInterval(c, "operatingTime")
	list, total, err := m.operLogService.List(c, &req)
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

func (m OperLogApi) Delete(c *gin.Context, req gdto.IDSReq) (any, error) {
	if err := m.operLogService.Delete(c, req.IDS); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}

func (m OperLogApi) DeleteAll(c *gin.Context, _ gdto.EmptyReq) (any, error) {
	if err := m.operLogService.DeleteAll(c); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}
