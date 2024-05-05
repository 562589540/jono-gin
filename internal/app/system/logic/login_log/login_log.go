package login_log

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"time"
)

var loginLogService service.ILoginLogService

type Service struct {
}

func New() service.ILoginLogService {
	if loginLogService == nil {
		loginLogService = &Service{}
	}
	return loginLogService
}

func (m *Service) Dao(ctx context.Context) dal.ILoginLogDo {
	return dal.LoginLog.WithContext(ctx)
}

func (m *Service) Create(ctx context.Context, param *dto.LoginParam) {
	ghub.Pool.Submit(func() {
		mModel := model.LoginLog{}
		mModel.Ip = param.Ip
		mModel.UserName = param.UserName
		mModel.Address = gutils.GetCityByIp(param.Ip)
		mModel.Behavior = param.Behavior
		mModel.Browser = param.Ua.Name
		mModel.Status = param.Status
		mModel.System = param.Ua.OS
		mModel.LoginTime = time.Now()
		err := m.Dao(ctx).Create(&mModel)
		gutils.CheckError(err)
	})
}

func (m *Service) Delete(ctx context.Context, ids []uint) error {
	dao := dal.LoginLog
	info, err := dao.WithContext(ctx).Where(dao.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf(constants.DeleteError)
	}
	return nil
}

func (m *Service) List(ctx context.Context, search *dto.LoginLogSearchReq) ([]dto.LoginLog, int64, error) {
	dp := dal.LoginLog
	q := dp.WithContext(ctx)

	if search.Status != "" {
		q = q.Where(dp.Status.Is(search.Status == "1"))
	}

	if search.UserName != "" {
		q = q.Where(dp.UserName.Eq(search.UserName))
	}

	if search.LoginTime != nil && len(search.LoginTime) == 2 {
		q = q.Where(dp.LoginTime.Between(search.LoginTime[0], search.LoginTime[1]))
	}

	list := make([]dto.LoginLog, 0)
	count, err := q.Order(dp.ID.Desc()).ScanByPage(&list, search.GetOffset(), search.GetLimit())
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (m *Service) DeleteAll(ctx context.Context) error {
	dao := dal.LoginLog
	_, err := dao.WithContext(ctx).Where(dao.ID.Gt(0)).Delete()
	if err != nil {
		return err
	}
	return nil
}
