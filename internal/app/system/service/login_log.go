package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/system/dto"
)

type ILoginLogService interface {
	Create(ctx context.Context, data *dto.LoginParam)
	Delete(ctx context.Context, ids []uint) error
	DeleteAll(ctx context.Context) error
	List(ctx context.Context, search *dto.LoginLogSearchReq) ([]dto.LoginLog, int64, error)
}
