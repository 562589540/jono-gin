package system

import (
	gdto2 "github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type UserOnlineApi struct {
	userOnlineService service.IUserOnlineService
}

func NewUserOnlineApi(userOnlineService service.IUserOnlineService) *UserOnlineApi {
	return &UserOnlineApi{
		userOnlineService: userOnlineService,
	}
}

func (m UserOnlineApi) List(c *gin.Context, mDto dto.UserOnlineSearchReq) (any, error) {
	list, total, err := m.userOnlineService.List(c, &mDto)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Data: gdto2.ListRes{
			List:  list,
			Total: total,
		},
	}, nil
}

func (m UserOnlineApi) Delete(c *gin.Context, mDto gdto2.IDReq) (any, error) {
	if err := m.userOnlineService.Delete(c, mDto.ID); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}
