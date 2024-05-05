package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
)

type IDeptService interface {
	Dao(ctx context.Context) dal.IDeptDo
	Create(ctx context.Context, data *dto.DeptAddReq) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, search *dto.DeptSearchReq) ([]dto.Dept, int64, error)
	Update(ctx context.Context, data *dto.DeptAddReq) error
}
