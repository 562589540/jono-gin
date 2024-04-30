package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type ITokenService interface {
	GenerateLoginToken(ctx context.Context, uid uint, userName string) (string, time.Time, error)
	GenerateRefreshToken(ctx context.Context, uid uint, userName string) (string, error)
	RefreshToken(ctx *gin.Context) (token string, refreshToken string, expireTime time.Time, err error)
}
