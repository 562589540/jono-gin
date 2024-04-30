package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
)

type IUserOnlineService interface {
	Dao(ctx context.Context) dal.IUserOnlineDo
	Create(ctx context.Context, data *dto.LoginParam)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, search *dto.UserOnlineSearchReq) ([]dto.UserOnline, int64, error)
}
