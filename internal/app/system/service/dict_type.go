package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
)

type IDictTypeService interface {
	Dao(ctx context.Context) dal.IDictTypeDo
	Create(ctx context.Context, data *dto.DictTypeAddReq) error
	Delete(ctx context.Context, ids []uint) error
	Update(ctx context.Context, data *dto.DictTypeUpdateReq) error
	List(ctx context.Context, search *dto.DictTypeSearchReq) ([]dto.DictType, int64, error)
	GetDictData(ctx context.Context, data *dto.DictGetReq) (*dto.DictGetRes, error)
	BatchGetDictData(ctx context.Context, data *dto.DictBatchGetReq) (map[string]*dto.DictGetRes, error)
}
