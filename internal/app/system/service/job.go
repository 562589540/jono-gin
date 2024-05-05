package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
)

type IJobService interface {
	Dao(ctx context.Context) dal.ISysJobDo
	Create(ctx context.Context, data *dto.JobAddReq) error
	DeleteJobLog(ctx context.Context, ids []uint) error
	DeleteJobLogAll(ctx context.Context, id uint) error
	Delete(ctx context.Context, ids []uint) error
	Once(ctx context.Context, id uint) error
	Update(ctx context.Context, data *dto.JobUpdateReq) error
	List(ctx context.Context, data *dto.JobSearchReq) ([]dto.Job, int64, error)
	JobLog(ctx context.Context, search *dto.JobLogReq) ([]dto.TaskLogRes, int64, error)
}
