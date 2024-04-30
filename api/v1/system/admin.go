package system

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type AdminApi struct {
	adminService service.IAdminService
}

func NewAdminApi(adminService service.IAdminService) *AdminApi {
	return &AdminApi{
		adminService: adminService,
	}
}

func (m AdminApi) List(c *gin.Context, req dto.AdminSearchReq) (any, error) {
	list, total, err := m.adminService.List(c, &req)
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

func (m AdminApi) Create(c *gin.Context, req dto.AdminAddReq) (any, error) {
	if err := m.adminService.Create(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "创建成功",
	}, nil
}

func (m AdminApi) Update(c *gin.Context, req dto.AdminUpdateReq) (any, error) {
	if err := m.adminService.Update(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m AdminApi) UpdateStatus(c *gin.Context, req dto.AdminUpdateStatusReq) (any, error) {
	if err := m.adminService.UpdateStatus(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m AdminApi) UpdatePassword(c *gin.Context, req dto.AdminUpdatePassReq) (any, error) {
	if err := m.adminService.UpdatePassword(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m AdminApi) UpdateAvatar(c *gin.Context, req dto.AdminUpdateAvatar) (any, error) {
	if err := m.adminService.UpdateAvatar(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m AdminApi) UpdateAdminRole(c *gin.Context, req dto.AdminUpdateRoleReq) (any, error) {
	if err := m.adminService.UpdateRole(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m AdminApi) GetAdminRoleIds(c *gin.Context, req gdto.IDReq) (any, error) {
	list, err := m.adminService.GetUserRoleIds(c, req.ID)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "查询成功",
		Data:    list,
	}, nil
}

func (m AdminApi) Delete(c *gin.Context, req gdto.IDSReq) (any, error) {
	if err := m.adminService.BatchDelete(c, req.IDS); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}
