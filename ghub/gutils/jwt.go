package gutils

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub/gbootstrap"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func getSigningKey() []byte {
	return []byte(gbootstrap.Cfg.Jwt.SigningKey)
}

type JwtCostClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(id uint, name string) (string, error) {
	jwtCostClaims := JwtCostClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			//token过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(gbootstrap.Cfg.Jwt.TokenExpire * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCostClaims)
	return token.SignedString(getSigningKey())
}

// GenerateRefreshToken 生成刷新令牌
func GenerateRefreshToken(id uint, name string) (string, error) {
	jwtCostClaims := JwtCostClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(gbootstrap.Cfg.Jwt.RefreshTokenExpire * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "RefreshToken",
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCostClaims)
	return refreshToken.SignedString(getSigningKey())
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*JwtCostClaims, error) {
	jwtCostClaims := &JwtCostClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, jwtCostClaims, func(token *jwt.Token) (interface{}, error) {
		return getSigningKey(), nil
	})
	if err == nil && !token.Valid {
		err = fmt.Errorf("无效的token")
	}
	return jwtCostClaims, err
}

// IsTokenValid 检查token是否有效
func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}
	return true
}

//	token, err := gutils.GenerateToken(1222, "234")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println(token)
//
//	parseToken, _ := gutils.ParseToken(token)
//	fmt.Println(parseToken)
