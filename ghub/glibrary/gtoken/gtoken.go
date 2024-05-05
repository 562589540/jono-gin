package gtoken

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/constants"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"time"
)

const (
	TokenName   = "Authorization"
	TokenPrefix = "Bearer "
)

type ResolveGinTokenOption struct {
	IsRefresh bool
	Token     string
}

type RefreshTokenJson struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// RedisTokenKey 获取redis token的key
func RedisTokenKey(uid uint) string {
	return strings.Replace(constants.LoginUserTokenRedisKey, "{ID}", strconv.Itoa(int(uid)), -1)
}

// RefreshRedisTokenKey 获取redis RefreshToken的key
func RefreshRedisTokenKey(uid uint) string {
	return strings.Replace(constants.RefreshUserTokenRedisKey, "{ID}", strconv.Itoa(int(uid)), -1)
}

// GenerateAndCacheLoginToken 生成token并且缓存
func GenerateAndCacheLoginToken(ctx context.Context, userId uint, userName string) (string, time.Time, error) {
	duration := ghub.Cfg.Jwt.TokenExpire * time.Minute
	expireTime := time.Now().Add(duration)
	token, err := gutils.GenerateToken(userId, userName)
	if err == nil {
		err = ghub.RedisClient.Set(ctx, RedisTokenKey(userId), token, duration)
	}
	return token, expireTime, err
}

// GenerateAndCacheRefreshToken 生成刷新token并且缓存
func GenerateAndCacheRefreshToken(ctx context.Context, userId uint, userName string) (string, error) {
	token, err := gutils.GenerateRefreshToken(userId, userName)
	if err == nil {
		err = ghub.RedisClient.Set(ctx, RefreshRedisTokenKey(userId), token, ghub.Cfg.Jwt.RefreshTokenExpire*time.Hour)
	}
	return token, err
}

// 创建并且缓存token
func generateAndCacheAllToken(ctx context.Context, userId uint, userName string) (token string, refreshToken string, expireTime time.Time, err error) {
	//生成登陆token
	token, expireTime, err = GenerateAndCacheLoginToken(ctx, userId, userName)
	if err != nil {
		return
	}

	//生成刷新token
	refreshToken, err = GenerateAndCacheRefreshToken(ctx, userId, userName)
	if err != nil {
		return
	}
	return
}

func ResolveGinToken(ctx *gin.Context, option ResolveGinTokenOption) (jwtCostClaims *gutils.JwtCostClaims, expireDuration time.Duration, err error) {
	var (
		rToken    string
		rTokenKey string
		tokenReq  = option.Token
	)

	//验证 在Header中获取
	if tokenReq == "" {
		tokenReq = ctx.GetHeader(TokenName)
		if tokenReq == "" || !strings.HasPrefix(tokenReq, TokenPrefix) {
			err = fmt.Errorf("缺少参数")
			return
		}
		tokenReq = tokenReq[len(TokenPrefix):]
	}

	//解析token
	jwtCostClaims, err = gutils.ParseToken(tokenReq)
	userId := jwtCostClaims.ID
	if err != nil || userId == 0 {
		ghub.Log.Error(err)
		err = fmt.Errorf("参数解析失败")
		return
	}

	if option.IsRefresh {
		rTokenKey = RefreshRedisTokenKey(userId)
	} else {
		rTokenKey = RedisTokenKey(userId)
	}
	rToken, err = ghub.RedisClient.Get(ctx, rTokenKey)
	if err != nil || tokenReq != rToken {
		err = fmt.Errorf("登陆状态已过期,请重新登陆")
		return
	}

	expireDuration, err = ghub.RedisClient.GetExpireDuration(ctx, rTokenKey)
	if err != nil || expireDuration <= 0 {
		err = fmt.Errorf("登陆状态已过期,请重新登陆")
		return
	}
	return
}

// RefreshToken 续签token
func RefreshToken(c *gin.Context) (token string, refreshToken string, expireTime time.Time, err error) {
	var (
		jwtCostClaims    *gutils.JwtCostClaims
		refreshTokenJson RefreshTokenJson
	)

	if err = c.BindJSON(&refreshTokenJson); err != nil {
		return
	}
	jwtCostClaims, _, err = ResolveGinToken(c, ResolveGinTokenOption{
		IsRefresh: true,
		Token:     refreshTokenJson.RefreshToken,
	})
	if err != nil {
		return
	}
	//生成token
	token, refreshToken, expireTime, err = generateAndCacheAllToken(c, jwtCostClaims.ID, jwtCostClaims.Name)
	if err != nil {
		return
	}
	return
}

// ClearAllLoginToken 清理用户的所有token
func ClearAllLoginToken(ctx context.Context, userId uint) error {
	rRefToken := RefreshRedisTokenKey(userId)
	rTokenKey := RedisTokenKey(userId)
	return ghub.RedisClient.ExecuteTransaction(ctx, func(pipe redis.Pipeliner) error {
		if err := pipe.Del(ctx, rRefToken, rTokenKey).Err(); err != nil {
			return err
		}
		return nil
	})
}

// CheckTokenIsExpire 查看token是否过期 检查的是刷新token
func CheckTokenIsExpire(ctx context.Context, userId uint) bool {
	_, err := ghub.RedisClient.Get(ctx, RefreshRedisTokenKey(userId))
	if err != nil {
		return true
	}
	return false
}
