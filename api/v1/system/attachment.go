package system

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type AttachmentApi struct {
	attachmentService service.IAttachmentService
}

func NewAttachmentApi(attachmentService service.IAttachmentService) *AttachmentApi {
	return &AttachmentApi{
		attachmentService: attachmentService,
	}
}

func (m AttachmentApi) List(c *gin.Context, req dto.AttachmentSearchReq) (any, error) {
	list, total, err := m.attachmentService.List(c, &req)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "查询成功",
		Data: gdto.ListRes{
			List:  list,
			Total: total,
		},
	}, nil
}

func (m AttachmentApi) Delete(c *gin.Context, req gdto.SnowflakeSReq) (any, error) {
	if err := m.attachmentService.Delete(c, req.Int64IDS()); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}
