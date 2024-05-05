package bizctx

import (
	"context"
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

func (m *Service) convertGinCtx(ctx context.Context) (*gin.Context, error) {
	if c, ok := ctx.(*gin.Context); ok {
		return c, nil
	}
	return nil, fmt.Errorf("上下文转换失败")
}

func (m *Service) GetLoginUser(ctx context.Context) (*model.LoginUser, error) {
	c, err := m.convertGinCtx(ctx)
	if err != nil {
		return nil, err
	}
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

func (m *Service) GetLoginUserModel(ctx context.Context) (*sysModel.Admin, error) {
	c, err := m.convertGinCtx(ctx)
	if err != nil {
		return nil, err
	}
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
