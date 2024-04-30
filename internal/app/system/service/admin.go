package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
)

type IAdminService interface {
	Dao(ctx context.Context) dal.IAdminDo
	SetLogin(ctx context.Context, userName, ip string)
	List(ctx context.Context, search *dto.AdminSearchReq) ([]dto.Admin, int64, error)
	Create(ctx context.Context, data *dto.AdminAddReq) error
	Delete(ctx context.Context, id uint) error
	BatchDelete(ctx context.Context, ids []uint) error
	Update(ctx context.Context, data *dto.AdminUpdateReq) error
	UpdateStatus(ctx context.Context, data *dto.AdminUpdateStatusReq) error
	UpdatePassword(ctx context.Context, data *dto.AdminUpdatePassReq) error
	UpdateAvatar(ctx context.Context, data *dto.AdminUpdateAvatar) error
	UpdateRole(ctx context.Context, data *dto.AdminUpdateRoleReq) error
	GetUserRoleIds(ctx context.Context, id uint) ([]uint, error)
}
