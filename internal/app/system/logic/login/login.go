package login

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/system/dal"
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
	ad := dal.Admin
	model, err = ad.WithContext(ctx).Preload(ad.RoleSign).Where(ad.UserName.Eq(data.UserName)).First()
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
	return gutils.RefreshToken(ctx)
}
