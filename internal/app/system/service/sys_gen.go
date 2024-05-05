package service

import (
	"context"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
)

type IGenService interface {
	Delete(ctx context.Context, ids []uint) error
	GetCodes(ctx context.Context, id uint) (*dto.GenCodeRes, error)
	GenCode(ctx context.Context, id uint) error
	GinInfo(ctx context.Context, req dto.TableInfoSearchReq) (*model.GenDate, error)
	List(ctx context.Context, req dto.TableInfoSearchReq) ([]dto.GenListRes, int64, error)
	TableList(ctx context.Context, req dto.TableInfoSearchReq) ([]dto.TableInfoRes, int64, error)
	TableInfo(ctx context.Context, req dto.TableInfoSearchReq) ([]model.TableColumn, error)
	TableDetails(ctx context.Context, req dto.TableInfoSearchReq) (*pkg.BaseInfo, error)
	GenTableFields(ctx context.Context, modeList []model.TableColumn) []*pkg.TableFields
	ImportDate(ctx context.Context, req model.GenDate) error
}
