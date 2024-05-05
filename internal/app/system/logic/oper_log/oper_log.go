package oper_log

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"strconv"
)

var operLogService service.IOperLogService

type Service struct{}

func New() service.IOperLogService {
	if operLogService == nil {
		operLogService = &Service{}
	}
	return operLogService
}

func (m *Service) Dao(ctx context.Context) dal.IOperLogDo {
	return dal.OperLog.WithContext(ctx)
}

func (m *Service) Create(ctx context.Context, data *dto.OperLogAddReq) error {
	err := m.Dao(ctx).Create(&model.OperLog{})
	if err != nil {
		return err
	}
	return nil
}

func (m *Service) Delete(ctx context.Context, ids []uint) error {
	info, err := m.Dao(ctx).Where(dal.OperLog.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf(constants.DeleteError)
	}
	return nil
}

func (m *Service) DeleteAll(ctx context.Context) error {
	dao := dal.OperLog
	_, err := dao.WithContext(ctx).Where(dao.ID.Gt(0)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (m *Service) Update(ctx context.Context, data *dto.OperLogUpdateReq) error {
	dp := dal.OperLog
	mModel, err := dp.WithContext(ctx).Where(dp.ID.Eq(data.ID)).First()
	if err != nil {
		return fmt.Errorf(constants.NoDataFound)
	}
	return dp.WithContext(ctx).Save(mModel)
}

func (m *Service) List(ctx context.Context, search *dto.OperLogSearchReq) ([]dto.OperLog, int64, error) {
	dp := dal.OperLog
	q := dp.WithContext(ctx)

	if search.Status != "" {
		status, err := strconv.Atoi(search.Status)
		if err == nil {
			q = q.Where(dp.Status.Eq(status))
		}
	}

	if search.Module != "" {
		q = q.Where(dp.Module.Eq(search.Module))
	}

	if search.CreatedAt != nil && len(search.CreatedAt) == 2 {
		q = q.Where(dp.CreatedAt.Between(search.CreatedAt[0], search.CreatedAt[1]))
	}

	list := make([]dto.OperLog, 0)
	count, err := q.Order(dp.ID.Desc()).ScanByPage(&list, search.GetOffset(), search.GetLimit())
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
