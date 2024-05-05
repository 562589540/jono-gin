package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/common/model"
	sysModel "github.com/562589540/jono-gin/internal/app/system/model"
)

type IContextService interface {
	GetLoginUser(ctx context.Context) (*model.LoginUser, error)
	GetLoginUserModel(ctx context.Context) (*sysModel.Admin, error)
}
