package bizctx

import (
	"fmt"
	"github.com/562589540/jono-gin/internal/app/common/model"
	sysModel "github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	instance *Service
	once     sync.Once
)

type Service struct {
}

func New() service.IContextService {
	once.Do(func() {
		instance = &Service{}
	})
	return instance
}

func (m *Service) GetLoginUser(c *gin.Context) (*model.LoginUser, error) {
	nUser, exists := c.Get(constants.LoginUser)
	if !exists {
		return nil, fmt.Errorf("为获取到上下文数据")
	}
	mUser, ok := nUser.(model.LoginUser)
	if !ok {
		return nil, fmt.Errorf("解析用户上下文数据失败")
	}
	return &mUser, nil
}

func (m *Service) GetLoginUserModel(c *gin.Context) (*sysModel.Admin, error) {
	uAdmin, exists := c.Get(constants.LoginAdminMode)
	if !exists {
		return nil, fmt.Errorf("为获取到上下文数据")
	}
	mAdmin, ok := uAdmin.(*sysModel.Admin)
	if !ok {
		return nil, fmt.Errorf("解析用户model上下文数据失败")
	}
	return mAdmin, nil
}
