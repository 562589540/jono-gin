package dictType

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"strconv"
)

var dictTypeService service.IDictTypeService

type Service struct{}

func New() service.IDictTypeService {
	if dictTypeService == nil {
		dictTypeService = &Service{}
	}
	return dictTypeService
}

func (m *Service) Create(ctx context.Context, data *dto.DictTypeAddReq) error {
	return dal.DictType.WithContext(ctx).Create(&model.DictType{
		DictName: data.DictName,
		DictType: data.DictType,
		Remark:   data.Remark,
		Status:   data.Status,
	})
}

func (m *Service) Delete(ctx context.Context, ids []uint) error {
	_, err := dal.DictType.WithContext(ctx).Where(dal.DictType.DictID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (m *Service) Update(ctx context.Context, data *dto.DictTypeUpdateReq) error {
	dp := dal.DictType
	mModel, err := dp.WithContext(ctx).Where(dp.DictID.Eq(data.ID)).First()
	if err != nil {
		return fmt.Errorf(constants.NoDataFound)
	}
	mModel.DictType = data.DictType
	mModel.DictName = data.DictName
	mModel.Status = data.Status
	mModel.Remark = data.Remark
	return dp.WithContext(ctx).Save(mModel)
}

func (m *Service) List(ctx context.Context, search *dto.DictTypeSearchReq) ([]dto.DictType, int64, error) {
	dt := dal.DictType
	q := dt.WithContext(ctx)

	if search.Status != "" {
		status, err := strconv.ParseInt(search.Status, 10, 32)
		if err != nil {
			return nil, 0, err
		}
		q = q.Where(dt.Status.Eq(int32(status)))
	}

	if search.DictName != "" {
		q = q.Where(dt.DictName.Eq(search.DictName))
	}

	if search.DictType != "" {
		q = q.Where(dt.DictType.Eq(search.DictType))
	}

	list := make([]dto.DictType, 0)
	count, err := q.ScanByPage(&list, search.GetOffset(), search.GetLimit())
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
