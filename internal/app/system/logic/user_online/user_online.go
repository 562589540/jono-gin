package user_online

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"time"
)

var userOnlineService service.IUserOnlineService

type Service struct{}

func New() service.IUserOnlineService {
	if userOnlineService == nil {
		userOnlineService = &Service{}
	}
	return userOnlineService
}

func (m *Service) Dao(ctx context.Context) dal.IUserOnlineDo {
	return dal.UserOnline.WithContext(ctx)
}

func (m *Service) Create(ctx context.Context, param *dto.LoginParam) {
	ghub.Pool.Submit(func() {
		uo := dal.UserOnline
		mModel, err := uo.WithContext(ctx).Where(uo.UserName.Eq(param.UserName)).FirstOrInit()
		if err != nil {
			ghub.ErrLog(err)
			return
		}
		mModel.Ip = param.Ip
		mModel.Uid = param.UserId
		mModel.System = param.Ua.OS
		mModel.Browser = param.Ua.Name
		mModel.Address = gutils.GetCityByIp(param.Ip)
		mModel.LoginTime = time.Now()
		err = uo.WithContext(ctx).Save(mModel)
		if err != nil {
			ghub.ErrLog(err)
		}
	})
}

func (m *Service) Delete(ctx context.Context, id uint) error {
	return dal.Q.Transaction(func(tx *dal.Query) error {
		uo := tx.UserOnline
		mModel, err := uo.WithContext(ctx).Where(uo.ID.Eq(id)).First()
		if err != nil {
			return err
		}
		_, err = uo.WithContext(ctx).Where(uo.ID.Eq(id)).Delete()
		if err != nil {
			return err
		}
		if err = gutils.ClearAllLoginToken(ctx, mModel.Uid); err != nil {
			return err
		}
		return nil
	})
}

func (m *Service) List(ctx context.Context, search *dto.UserOnlineSearchReq) ([]dto.UserOnline, int64, error) {
	m.clearUserOnlineExpireData(ctx)
	uo := dal.UserOnline
	q := uo.WithContext(ctx)

	if search.UserName != "" {
		q = q.Where(uo.UserName.Like(fmt.Sprintf("%%%s%%", search.UserName)))
	}

	s := make([]dto.UserOnline, 0)
	total, err := q.Order(uo.ID.Desc()).ScanByPage(&s, search.GetOffset(), search.GetLimit())
	if err != nil {
		return nil, 0, err
	}
	return s, total, nil
}

func (m *Service) clearUserOnlineExpireData(ctx context.Context) {
	// 使用批量方式处理过期Token
	uo := dal.UserOnline
	mModels, err := uo.WithContext(ctx).Find()
	if err != nil {
		return
	}
	var toDelete []uint
	for _, mModel := range mModels {
		if gutils.CheckTokenIsExpire(ctx, mModel.Uid) {
			toDelete = append(toDelete, mModel.ID)
		}
	}
	//清理过期的数据
	if len(toDelete) > 0 {
		_, err = uo.WithContext(ctx).Where(uo.ID.In(toDelete...)).Delete()
		ghub.ErrLog(err)
	}
}
