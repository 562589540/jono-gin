package user_online

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/glibrary/gtoken"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
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
		dao := dal.UserOnline
		mModel, err := dao.WithContext(ctx).Where(dao.UserName.Eq(param.UserName)).FirstOrInit()
		if err != nil {
			gutils.Error(err)
			return
		}
		mModel.Ip = param.Ip
		mModel.Uid = param.UserId
		mModel.System = param.Ua.OS
		mModel.Browser = param.Ua.Name
		mModel.Address = gutils.GetCityByIp(param.Ip)
		mModel.LoginTime = time.Now()
		err = dao.WithContext(ctx).Save(mModel)
		gutils.CheckError(err)
	})
}

func (m *Service) Delete(ctx context.Context, id uint) error {
	return dal.Q.Transaction(func(tx *dal.Query) error {
		dao := tx.UserOnline
		mModel, err := dao.WithContext(ctx).Where(dao.ID.Eq(id)).First()
		if err != nil {
			return err
		}
		info, err := dao.WithContext(ctx).Where(dao.ID.Eq(id)).Delete()
		if err != nil {
			return err
		}
		if info.RowsAffected == 0 {
			return fmt.Errorf(constants.DeleteError)
		}
		if err = gtoken.ClearAllLoginToken(ctx, mModel.Uid); err != nil {
			return err
		}
		return nil
	})
}

func (m *Service) List(ctx context.Context, search *dto.UserOnlineSearchReq) ([]dto.UserOnline, int64, error) {
	m.clearUserOnlineExpireData(ctx)
	dao := dal.UserOnline
	q := dao.WithContext(ctx)

	if search.UserName != "" {
		q = q.Where(dao.UserName.Like(fmt.Sprintf("%%%s%%", search.UserName)))
	}

	s := make([]dto.UserOnline, 0)
	total, err := q.Order(dao.ID.Desc()).ScanByPage(&s, search.GetOffset(), search.GetLimit())
	if err != nil {
		return nil, 0, err
	}
	return s, total, nil
}

func (m *Service) clearUserOnlineExpireData(ctx context.Context) {
	// 使用批量方式处理过期Token
	dao := dal.UserOnline
	mModels, err := dao.WithContext(ctx).Find()
	if err != nil {
		return
	}
	var toDelete []uint
	for _, mModel := range mModels {
		if gtoken.CheckTokenIsExpire(ctx, mModel.Uid) {
			toDelete = append(toDelete, mModel.ID)
		}
	}
	//清理过期的数据
	if len(toDelete) > 0 {
		_, err = dao.WithContext(ctx).Where(dao.ID.In(toDelete...)).Delete()
		gutils.CheckError(err)
	}
}
