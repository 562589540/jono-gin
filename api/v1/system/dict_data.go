package system

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type DictDataApi struct {
	dictDataService service.IDictDataService
}

func NewDictDataApi(dictDataService service.IDictDataService) *DictDataApi {
	return &DictDataApi{
		dictDataService: dictDataService,
	}
}

func (m DictDataApi) Create(c *gin.Context, req dto.DictDataAddReq) (any, error) {
	if err := m.dictDataService.Create(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "创建成功",
	}, nil
}

func (m DictDataApi) List(c *gin.Context, req dto.DictDataSearchReq) (any, error) {
	list, total, err := m.dictDataService.List(c, &req)
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

func (m DictDataApi) Update(c *gin.Context, req dto.DictDataUpdateReq) (any, error) {
	if err := m.dictDataService.Update(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m DictDataApi) Delete(c *gin.Context, req gdto.IDSReq) (any, error) {
	if err := m.dictDataService.Delete(c, req.IDS); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}
