package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/gin-gonic/gin"
)

type IMenuService interface {
	Dao(ctx context.Context) dal.IMenuDo
	Create(ctx context.Context, data *dto.AddMenuReq) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]dto.Menu, error)
	Update(ctx context.Context, data *dto.UpdateMenuReq) error
	GetRoleMenu(ctx context.Context) ([]dto.RoleMenu, error)
	GetRoutes(ctx *gin.Context) ([]*dto.MenuList, error)
}
