package middleware

import (
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/logic/bizctx"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
	"strings"
)

func PostRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//startTime := time.Now()

		// 请求处理函数之前没有任何代码，先让路由处理器运行
		//fmt.Println("-----------------------前置-------------------------------")
		c.Next()
		//duration := time.Since(startTime)
		//fmt.Println("-----------------------后置-------------------------------")
		// 这里的代码将在请求处理完毕后执行
		// 获取状态码等信息进行日志记录
		statusCode := c.Writer.Status()

		//log.Printf("After request | status: %d | method: %s | path: %s | duration: %s",
		//	statusCode, c.Request.Method, c.Request.URL.Path, duration)

		//2024/04/23 18:03:33 After request | status: 200 | method: GET | path: /api/v1/system/menu/getRoutes | duration: 6.74625ms

		ghub.Pool.Submit(func() {
			user, err := bizctx.New().GetLoginUser(c)
			if err != nil {
				return
			}
			ua := useragent.Parse(c.GetHeader("User-Agent"))
			ip := c.ClientIP()
			s := 0
			if statusCode == 200 {
				s = 1
			}
			err = dal.OperLog.Create(&model.OperLog{
				Ip:       ip,
				Browser:  ua.Name,
				System:   ua.OS,
				UserName: user.UserName,
				Address:  gutils.GetCityByIp(ip),
				Module:   extract(c.Request.URL.Path, 2),
				Summary:  extract(c.Request.URL.Path, 1),
				Status:   s,
			})
			ghub.ErrLog(err)
		})
	}
}

func extract(url string, index int) string {
	parts := strings.Split(url, "/")
	l := len(parts)
	if l < 2 {
		return ""
	}
	i := len(parts) - index
	if i < 0 {
		return ""
	}
	penultimate := parts[i]
	return penultimate
}
