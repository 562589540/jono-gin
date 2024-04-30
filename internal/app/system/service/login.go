package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
)

type ILoginService interface {
	Login(ctx context.Context, req *dto.AdminLoginReq) (model *model.Admin, err error)
}
