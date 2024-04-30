package system

import (
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type MysqlApi struct {
	service service.IGenService
}

func NewGenApi(service service.IGenService) *MysqlApi {
	return &MysqlApi{
		service: service,
	}
}

func (m MysqlApi) List(c *gin.Context, req dto.TableInfoSearchReq) (any, error) {
	req.Time, _ = ghub.ParseTimeInterval(c, "createTime")
	list, total, err := m.service.List(c, req)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Data: gdto.ListRes{
			List:  list,
			Total: total,
		},
	}, nil
}

func (m MysqlApi) TableList(c *gin.Context, req dto.TableInfoSearchReq) (any, error) {
	req.Time, _ = ghub.ParseTimeInterval(c, "createTime")
	list, total, err := m.service.TableList(c, req)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Data: gdto.ListRes{
			List:  list,
			Total: total,
		},
	}, nil
}

func (m MysqlApi) TableInfo(c *gin.Context, req dto.TableInfoSearchReq) (any, error) {
	genDate, err := m.service.GinInfo(c, req)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Data: gin.H{
			"details": genDate,
		},
	}, nil
}

func (m MysqlApi) ImportDate(c *gin.Context, req model.GenDate) (any, error) {
	if err := m.service.ImportDate(c, req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "导入成功",
	}, nil
}

func (m MysqlApi) GenCode(c *gin.Context, req gdto.IDReq) (any, error) {
	if err := m.service.GenCode(c, req.ID); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "代码生成成功",
	}, nil
}

func (m MysqlApi) GetCode(c *gin.Context, req gdto.IDReq) (any, error) {
	codes, err := m.service.GetCodes(c, req.ID)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "获取成功",
		Data:    codes,
	}, nil
}

func (m MysqlApi) Delete(c *gin.Context, req gdto.IDSReq) (any, error) {
	if err := m.service.Delete(c, req.IDS); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}
