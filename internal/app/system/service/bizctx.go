package service

import (
	"github.com/562589540/jono-gin/internal/app/common/model"
	sysModel "github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/gin-gonic/gin"
)

type IContextService interface {
	GetLoginUser(c *gin.Context) (*model.LoginUser, error)
	GetLoginUserModel(c *gin.Context) (*sysModel.Admin, error)
}
