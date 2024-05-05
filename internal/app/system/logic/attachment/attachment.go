package attachment

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants/enum"
	"os"
)

var attachmentService service.IAttachmentService

type Service struct{}

func New() service.IAttachmentService {
	if attachmentService == nil {
		attachmentService = &Service{}
	}
	return attachmentService
}

func (m *Service) Dao(ctx context.Context) dal.IAttachmentDo {
	return dal.Attachment.WithContext(ctx)
}

func (m *Service) Delete(ctx context.Context, ids []int64) error {
	fmt.Println(ids)

	dao := dal.Attachment
	attachmentList, err := dao.WithContext(ctx).Where(dao.ID.In(ids...)).Find()
	if err != nil {
		return err
	}
	var successCount int
	fmt.Println(attachmentList)

	for _, attachment := range attachmentList {
		//保证删除文件与数据同步使用事务
		gutils.CheckError(dal.Q.Transaction(func(tx *dal.Query) error {
			if _, err = tx.Attachment.WithContext(ctx).Where(tx.Attachment.ID.Eq(attachment.ID)).Delete(); err != nil {
				gutils.CheckError(err)
				return err
			}
			//本地储存
			if attachment.StorageEngine == enum.Local {
				//有错误 文件并且存在报错 文件不存在算成功
				if err = os.Remove(attachment.Path); err != nil && !os.IsNotExist(err) {
					gutils.CheckError(err)
					return err
				}
			}
			successCount++
			return nil
		}))
	}
	//一个都没成功
	if successCount == 0 {
		return fmt.Errorf("删除失败2")
	}
	//有成功即可返回成功
	return nil
}

func (m *Service) List(ctx context.Context, search *dto.AttachmentSearchReq) ([]dto.Attachment, int64, error) {
	dao := dal.Attachment
	q := dao.WithContext(ctx).Where(dao.Status.Eq(1))
	//if search.FileName != "" {
	//	q = q.Where(dao.FileName.Eq(search.FileName))
	//}
	//if search.FileType != "" {
	//	q = q.Where(dao.FileType.Eq(search.FileType))
	//}
	//if search.StorageEngine != "" {
	//	q = q.Where(dao.StorageEngine.Eq(search.StorageEngine))
	//}
	//if search.Class != 0 {
	//	q = q.Where(dao.Class.Eq(search.Class))
	//}
	//if search.CreatedBy != 0 {
	//	q = q.Where(dao.CreatedBy.Eq(search.CreatedBy))
	//}
	list := make([]dto.Attachment, 0)
	count, err := q.Order(dao.ID.Desc()).ScanByPage(&list, search.GetOffset(), search.GetLimit())

	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
