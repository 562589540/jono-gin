package middleware

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/gbootstrap"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/logic/bizctx"
	"github.com/562589540/jono-gin/internal/app/system/logic/casbin"
	"github.com/562589540/jono-gin/internal/constants"
	"github.com/gin-gonic/gin"
)

func Init() {
	ghub.EventBus.Subscribe(constants.RefreshPathToRoles, func(data interface{}) {
		//注册刷新策略
		casbin.New().RefreshCasbinRule()
	})
}

func ruleErr(c *gin.Context, msg string) {
	gres.Ok(c, gres.Response{
		Code:    10401,
		Message: msg,
	})
}
func AuthRule() func(c *gin.Context) {
	return func(c *gin.Context) {
		mUser, err := bizctx.New().GetLoginUser(c)
		if err != nil {
			ruleErr(c, "无访问权限")
			c.Abort()
			return
		}

		ad := dal.Admin
		mAdminModel, err := ad.WithContext(c).Preload(ad.RoleSign).Where(dal.Admin.ID.Eq(mUser.ID)).First()
		if err != nil {
			ruleErr(c, "无访问权限3")
			c.Abort()
			return
		}
		if !ghub.Contains(gbootstrap.Cfg.System.NotCheckAuthAdminIds, mUser.ID) {
			hasRole := false
			path := c.Request.URL.Path
			for _, role := range mAdminModel.RoleSign {
				sub := fmt.Sprintf("u_%d", role.ID)
				obj := path
				act := "All"
				ok, _ := casbin.New().Enforcer.Enforce(sub, obj, act)
				// 用户拥有至少一个所需的角色，可以访问
				if ok == true {
					hasRole = true
					break
				}
			}
			if !hasRole {
				ruleErr(c, "无访问权限4")
				c.Abort()
				return
			}
		}

		c.Set(constants.LoginAdminMode, mAdminModel)
		c.Next()
	}
}
