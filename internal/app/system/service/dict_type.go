package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/system/dto"
)

type IDictTypeService interface {
	Create(ctx context.Context, data *dto.DictTypeAddReq) error
	Delete(ctx context.Context, ids []uint) error
	Update(ctx context.Context, data *dto.DictTypeUpdateReq) error
	List(ctx context.Context, search *dto.DictTypeSearchReq) ([]dto.DictType, int64, error)
}
