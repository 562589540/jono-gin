package uploader

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub/glibrary/gfile"
	"github.com/562589540/jono-gin/ghub/glibrary/gsnowflake"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/logic/bizctx"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants/enum"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	instance service.IUploaderService
	once     sync.Once
)

func New(localPath string, nodeNumber int64) service.IUploaderService {
	once.Do(func() {
		instance = &Service{
			localPath:  localPath,
			nodeNumber: nodeNumber,
		}
	})
	return instance
}

type Service struct {
	localPath  string
	nodeNumber int64
}

func (m *Service) Create(ctx context.Context, data dto.ChunkInfoReq) {
	//文件信息没有就创建
	err := dal.Attachment.WithContext(ctx).Create(&model.Attachment{
		ID:            gsnowflake.GetInstance(m.nodeNumber).Generate(),
		Hash:          data.Identifier,
		TotalChunks:   data.TotalChunks,
		StorageEngine: enum.Local,
		Status:        0,
	})
	gutils.CheckError(err)
}

// CheckChunkInfo 查看文件上传状态
func (m *Service) CheckChunkInfo(ctx context.Context, data dto.ChunkInfoReq) *dto.ChunkInfoRes {

	fileDao := dal.Attachment
	//查找文件
	first, err := fileDao.WithContext(ctx).Where(fileDao.Hash.Eq(data.Identifier)).First()
	if err != nil {
		//数据库没文件数据 创建数据
		m.Create(ctx, data)
		return &dto.ChunkInfoRes{
			SkipUpload: false,
		}
	}
	//完成状态检查文件是否存在
	if first.Status == 1 {
		//查看文件是否存在
		if m.checkFilesIsExist(ctx, first) {
			//文件存在直接返回秒传即可
			return &dto.ChunkInfoRes{
				SkipUpload: true,
			}
		}
	}

	chunkDao := dal.Chunk
	//查找分片
	find, err := chunkDao.WithContext(ctx).Where(chunkDao.Hash.Eq(first.Hash)).Find()
	if err != nil {
		return &dto.ChunkInfoRes{
			SkipUpload: false,
		}
	}
	var chunks []int
	for _, chunk := range find {
		if m.checkFilesChunkIsExist(ctx, chunk) {
			chunks = append(chunks, chunk.Number)
		}
	}
	return &dto.ChunkInfoRes{
		SkipUpload:     false,  //文件是否完成上传并且合并
		UploadedChunks: chunks, //上传成功的分片id
	}
}

// PutChunk 分片上传
func (m *Service) PutChunk(c *gin.Context, formFile string, req dto.ChunkInfoReq) error {
	return dal.Q.Transaction(func(tx *dal.Query) error {
		// 尝试从请求中获取分片文件
		file, err := c.FormFile(formFile)
		if err != nil {
			return err
		}

		// 获取文件扩展名并检查是否允许
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !isAllowedExtension(ext) {
			return fmt.Errorf("file type %s not allowed", ext)
		}

		// 存储分片的临时目录
		dir := filepath.Join(m.localPath, req.Identifier)
		// 保存文件分片
		filePath := filepath.Join(dir, m.convChunkNumber(req.ChunkNumber))

		//分片数据入库
		mModel, err := tx.Chunk.WithContext(c).Where(tx.Chunk.Hash.Eq(req.Identifier), tx.Chunk.Number.Eq(req.ChunkNumber)).FirstOrInit()
		if err != nil {
			return err
		}
		if mModel.ID == 0 {
			mModel.ID = gsnowflake.GetInstance(m.nodeNumber).Generate()
		}
		mModel.Number = req.ChunkNumber
		mModel.Path = dir
		mModel.Size = req.ChunkSize
		mModel.Hash = req.Identifier
		mModel.Status = 1
		if err = tx.Chunk.WithContext(c).Save(mModel); err != nil {
			return err
		}

		//创建目录
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
		//写入分片文件
		if err = c.SaveUploadedFile(file, filePath); err != nil {
			return err
		}
		return nil
	})
}

// MergeFile 合并分片
func (m *Service) MergeFile(ctx context.Context, data dto.MergeFileReq) error {
	// 检查文件类型
	ext := strings.ToLower(filepath.Ext(data.Name))
	if !isAllowedExtension(ext) {
		return fmt.Errorf("file type %s not allowed", ext)
	}

	fileDao := dal.Attachment
	//查找文件
	first, err := fileDao.WithContext(ctx).Where(fileDao.Hash.Eq(data.UniqueIdentifier)).First()
	if err != nil {
		return err
	}
	//完成状态检查文件是否存在
	if first.Status == 1 {
		//文件存在直接秒传成功
		if m.checkFilesIsExist(ctx, first) {
			_, _ = dal.Chunk.WithContext(ctx).Where(dal.Chunk.Hash.Eq(data.UniqueIdentifier)).Delete()
			return nil
		}
	}

	dir := filepath.Join(m.localPath, data.UniqueIdentifier)
	finalPath := filepath.Join(m.localPath, data.Name)

	if err = m.mergeFiles(dir, finalPath, first.TotalChunks); err != nil {
		return err
	}
	var createdId uint
	//更新
	userModel, err := bizctx.New().GetLoginUserModel(ctx)
	if err != nil || userModel == nil {
		createdId = 0
	} else {
		createdId = userModel.ID
	}
	_, _ = fileDao.WithContext(ctx).Where(fileDao.ID.Eq(first.ID)).Updates(model.Attachment{
		Status:    1,
		FileName:  data.Name,
		FileSize:  gfile.FormatBytes(int64(data.Size)),
		FileType:  data.FileType,
		Path:      finalPath,
		CreatedBy: createdId,
		Class:     0,
	})
	_, _ = dal.Chunk.WithContext(ctx).Where(dal.Chunk.Hash.Eq(data.UniqueIdentifier)).Delete()
	return nil
}

func (m *Service) mergeFiles(dir, finalPath string, totalChunks int) error {
	finalFile, err := os.Create(finalPath)
	if err != nil {
		return err
	}
	defer func() {
		gutils.CheckError(finalFile.Close())
		if err != nil { // 如果发生错误，删除不完整的最终文件
			gutils.CheckError(os.Remove(finalPath))
		}
	}()

	for i := 1; i <= totalChunks; i++ {
		chunkPath := filepath.Join(dir, m.convChunkNumber(i))
		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(finalFile, chunkFile)
		_ = chunkFile.Close() // 确保每个分片文件在复制后都被关闭
		if err != nil {
			return err
		}
	}
	// 合并成功后清理目录
	gutils.CheckError(os.RemoveAll(dir))
	return nil
}

// 检查文件是否存在 不存在重置状态
func (m *Service) checkFilesIsExist(ctx context.Context, attachment *model.Attachment) bool {
	if !gfile.FileIsExisted(attachment.Path) {
		_, err := dal.Attachment.WithContext(ctx).Where(dal.Attachment.ID.Eq(attachment.ID)).Update(dal.Attachment.Status, 0)
		gutils.CheckError(err)
		return false
	}
	return true
}

// 检查切片文件是否存在 不存在直接删除
func (m *Service) checkFilesChunkIsExist(ctx context.Context, chunk *model.Chunk) bool {
	if !gfile.FileIsExisted(chunk.Path) {
		_, err := dal.Chunk.WithContext(ctx).Where(dal.Chunk.ID.Eq(chunk.ID)).Delete()
		gutils.CheckError(err)
		return false
	}
	return true
}

func (m *Service) convChunkNumber(chunkNumber int) string {
	return fmt.Sprintf("chunk_%d", chunkNumber)
}

var allowedExtensions = map[string]bool{
	".png": true, ".jpg": true, ".jpeg": true, ".gif": true, ".bmp": true,
	".mp4": true, ".rmvb": true, ".mkv": true, ".wmv": true, ".flv": true,
	".doc": true, ".docx": true, ".xls": true, ".xlsx": true, ".ppt": true, ".pptx": true, ".pdf": true,
	".txt": true, ".tif": true, ".tiff": true, ".rar": true, ".zip": true,
}

func isAllowedExtension(extension string) bool {
	return allowedExtensions[strings.ToLower(extension)]
}
