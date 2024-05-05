package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
)

type IAttachmentService interface {
	Dao(ctx context.Context) dal.IAttachmentDo
	Delete(ctx context.Context, ids []int64) error
	List(ctx context.Context, data *dto.AttachmentSearchReq) ([]dto.Attachment, int64, error)
}
