package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
)

type IRolesService interface {
	Dao(ctx context.Context) dal.IRolesDo
	Create(ctx context.Context, data *dto.RolesAddReq) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, search *dto.RolesSearchReq) ([]dto.Role, int64, error)
	Update(ctx context.Context, data *dto.RolesUpdateReq) error
	UpdateRoleMenusAuth(ctx context.Context, data *dto.RolesPowerReq) error
	GetRoleMenuIds(ctx context.Context, id uint) ([]uint, error)
	GetAllRoleList(ctx context.Context) ([]dto.Role, error)
}
