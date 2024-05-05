package login

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub/glibrary/gtoken"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"github.com/gin-gonic/gin"
	"time"
)

var instance service.ILoginService

func New() service.ILoginService {
	if instance == nil {
		instance = &Service{}
	}
	return instance
}

type Service struct {
	tokenService service.ITokenService
}

func (m *Service) Login(ctx context.Context, data *dto.AdminLoginReq) (model *model.Admin, err error) {
	adminDao := dal.Admin
	model, err = adminDao.WithContext(ctx).Preload(adminDao.RoleSign).Where(adminDao.UserName.Eq(data.UserName)).First()
	if err != nil {
		err = fmt.Errorf(constants.UserNameError)
		return
	}
	if !gutils.CompareHashAndPassword(model.Password, data.Password) {
		err = fmt.Errorf(constants.PasswordError)
		return
	}
	return
}

// RefreshToken 续签token
func (m *Service) RefreshToken(ctx *gin.Context) (token string, refreshToken string, expireTime time.Time, err error) {
	return gtoken.RefreshToken(ctx)
}
