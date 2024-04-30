package system

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type DictTypeApi struct {
	dictTypeService service.IDictTypeService
}

func NewDictTypeApi(dictTypeService service.IDictTypeService) *DictTypeApi {
	return &DictTypeApi{
		dictTypeService: dictTypeService,
	}
}

func (m DictTypeApi) Create(c *gin.Context, req dto.DictTypeAddReq) (any, error) {
	if err := m.dictTypeService.Create(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "创建成功",
	}, nil
}

func (m DictTypeApi) List(c *gin.Context, req dto.DictTypeSearchReq) (any, error) {
	list, total, err := m.dictTypeService.List(c, &req)
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

func (m DictTypeApi) Update(c *gin.Context, req dto.DictTypeUpdateReq) (any, error) {
	if err := m.dictTypeService.Update(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m DictTypeApi) Delete(c *gin.Context, req gdto.IDSReq) (any, error) {
	if err := m.dictTypeService.Delete(c, req.IDS); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}
