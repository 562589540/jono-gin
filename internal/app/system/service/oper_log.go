package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
)

type IOperLogService interface {
	Dao(ctx context.Context) dal.IOperLogDo
	Create(ctx context.Context, data *dto.OperLogAddReq) error
	Delete(ctx context.Context, ids []uint) error
	DeleteAll(ctx context.Context) error
	Update(ctx context.Context, data *dto.OperLogUpdateReq) error
	List(ctx context.Context, search *dto.OperLogSearchReq) ([]dto.OperLog, int64, error)
}
