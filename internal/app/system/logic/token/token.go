package token

import (
	"context"
	"github.com/562589540/jono-gin/ghub/glibrary/gtoken"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
	"time"
)

var instance service.ITokenService

func New() service.ITokenService {
	if instance == nil {
		instance = &Service{}
	}
	return instance
}

type Service struct{}

func (s *Service) GenerateLoginToken(ctx context.Context, uid uint, userName string) (string, time.Time, error) {
	return gtoken.GenerateAndCacheLoginToken(ctx, uid, userName)
}

func (s *Service) GenerateRefreshToken(ctx context.Context, uid uint, userName string) (string, error) {
	return gtoken.GenerateAndCacheRefreshToken(ctx, uid, userName)
}

// RefreshToken 续签token
func (s *Service) RefreshToken(ctx *gin.Context) (token string, refreshToken string, expireTime time.Time, err error) {
	return gtoken.RefreshToken(ctx)
}
