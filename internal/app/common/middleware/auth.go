package middleware

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/common/model"
	"github.com/562589540/jono-gin/internal/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

func tokenErr(c *gin.Context, msg string) {
	gres.Fail(c, gres.Response{
		Status:  http.StatusUnauthorized,
		Code:    401,
		Message: msg,
	})
}

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		jwtCostClaims, _, err := gutils.ResolveGinToken(c, gutils.ResolveGinTokenOption{})
		if err != nil {
			tokenErr(c, err.Error())
			return
		}
		c.Set(constants.LoginUser, model.LoginUser{
			ID:       jwtCostClaims.ID,
			UserName: jwtCostClaims.Name,
		})
		c.Next()
	}
}

//后端续期token行为
//if expireDuration.Seconds() < (10 * time.Minute).Seconds() {
//	newToken, _, err := gutils.GenerateAndCacheLoginToken(userId, jwtCostClaims.NameZh)
//	if err != nil {
//		tokenErr(c)
//		return
//	}
//	c.Header("token", newToken)
//}

//用户信息保存到上下文
//user, err := dao.NewUserDao().GetUserByID(userId)
//if err != nil {
//	tokenErr(c)
//	return
//}
//c.Set(constants.LoginUser, user)
