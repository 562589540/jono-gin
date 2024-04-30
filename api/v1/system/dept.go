package system

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type DeptApi struct {
	deptService service.IDeptService
}

func NewDeptApi(deptService service.IDeptService) *DeptApi {
	return &DeptApi{
		deptService: deptService,
	}
}

func (m DeptApi) Create(c *gin.Context, req dto.DeptAddReq) (any, error) {
	if err := m.deptService.Create(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "创建成功",
	}, nil
}

func (m DeptApi) List(c *gin.Context, req dto.DeptSearchReq) (any, error) {
	list, _, err := m.deptService.List(c, &req)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Data: list,
	}, nil
}

func (m DeptApi) Update(c *gin.Context, req dto.DeptAddReq) (any, error) {
	if err := m.deptService.Update(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m DeptApi) Delete(c *gin.Context, req gdto.IDReq) (any, error) {
	if err := m.deptService.Delete(c, req.ID); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}
