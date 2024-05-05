package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
)

type IDictDataService interface {
	Dao(ctx context.Context) dal.IDictDataDo
	Create(ctx context.Context, data *dto.DictDataAddReq) error
	Delete(ctx context.Context, ids []uint) error
	Update(ctx context.Context, data *dto.DictDataUpdateReq) error
	List(ctx context.Context, data *dto.DictDataSearchReq) ([]dto.DictData, int64, error)
}
