package middleware

import (
	"github.com/gin-gonic/gin"
)

// UploaderAuth 上传限制等中间件
// 这里可以过滤ip 限制上传时间间隔等
func UploaderAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
	}
}
