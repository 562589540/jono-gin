package dictType

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"strconv"
)

var dictTypeService service.IDictTypeService

type Service struct {
	contextService service.IContextService
}

func New(contextService service.IContextService) service.IDictTypeService {
	if dictTypeService == nil {
		dictTypeService = &Service{
			contextService,
		}
	}
	return dictTypeService
}

func (m *Service) Dao(ctx context.Context) dal.IDictTypeDo {
	return dal.DictType.WithContext(ctx)
}

func (m *Service) Create(ctx context.Context, data *dto.DictTypeAddReq) error {
	mModel := new(model.DictType)
	if err := gutils.Copy(mModel, data); err != nil {
		return err
	}
	userModel, err := m.contextService.GetLoginUserModel(ctx)
	if err == nil {
		mModel.CreateBy = int32(userModel.ID)
	}
	return m.Dao(ctx).Create(mModel)
}

func (m *Service) Delete(ctx context.Context, ids []uint) error {
	dao := dal.DictType
	_, err := dao.WithContext(ctx).Where(dao.DictID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (m *Service) Update(ctx context.Context, data *dto.DictTypeUpdateReq) error {
	dao := dal.DictType
	mModel, err := dao.WithContext(ctx).Where(dao.DictID.Eq(data.ID)).First()
	if err != nil {
		return fmt.Errorf(constants.NoDataFound)
	}
	if err = gutils.Copy(mModel, data); err != nil {
		return err
	}
	userModel, err := m.contextService.GetLoginUserModel(ctx)
	if err == nil {
		mModel.UpdateBy = int32(userModel.ID)
	}
	return dao.WithContext(ctx).Save(mModel)
}

func (m *Service) List(ctx context.Context, search *dto.DictTypeSearchReq) ([]dto.DictType, int64, error) {
	dao := dal.DictType
	q := dao.WithContext(ctx)

	if search.Status != "" {
		status, err := strconv.ParseInt(search.Status, 10, 32)
		if err != nil {
			return nil, 0, err
		}
		q = q.Where(dao.Status.Eq(int32(status)))
	}

	if search.DictName != "" {
		q = q.Where(dao.DictName.Eq(search.DictName))
	}

	if search.DictType != "" {
		q = q.Where(dao.DictType.Eq(search.DictType))
	}

	list := make([]dto.DictType, 0)
	count, err := q.ScanByPage(&list, search.GetOffset(), search.GetLimit())
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (m *Service) GetDictData(ctx context.Context, data *dto.DictGetReq) (*dto.DictGetRes, error) {
	dao := dal.DictType
	d := dal.DictData
	mModel, err := dao.WithContext(ctx).Preload(dao.DictData.Order(d.DictSort.Asc(), d.DictValue)).Where(dao.DictType.Eq(data.DictType)).First()
	if err != nil {
		return nil, err
	}
	values := make([]*dto.DictDataValue, len(mModel.DictData))
	for i, datum := range mModel.DictData {
		values[i] = &dto.DictDataValue{
			Remark:    datum.Remark,
			IsDefault: datum.GetIsDefault(),
			Key:       datum.DictValue,
			Value:     datum.DictLabel,
		}
	}
	return &dto.DictGetRes{
		Info: &dto.DictDataInfo{
			Name:   mModel.DictName,
			Remark: mModel.Remark,
		},
		Values: values,
	}, nil
}

func (m *Service) BatchGetDictData(ctx context.Context, data *dto.DictBatchGetReq) (map[string]*dto.DictGetRes, error) {
	dao := dal.DictType
	mList, err := dao.WithContext(ctx).Preload(dao.DictData).Where(dao.DictType.In(data.DictTypes...)).Find()
	if err != nil {
		return nil, err
	}

	//批量字典列表
	dictList := make(map[string]*dto.DictGetRes)

	for _, dict := range mList {
		//单个字典
		values := make([]*dto.DictDataValue, len(dict.DictData))
		for i, datum := range dict.DictData {
			values[i] = &dto.DictDataValue{
				Remark:    datum.Remark,
				IsDefault: datum.GetIsDefault(),
				Key:       datum.DictValue,
				Value:     datum.DictLabel,
			}
		}
		dictList[dict.DictType] = &dto.DictGetRes{
			Info: &dto.DictDataInfo{
				Name:   dict.DictName,
				Remark: dict.Remark,
			},
			Values: values,
		}
	}

	return dictList, nil
}
