package system

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type UploaderApi struct {
	service   service.IUploaderService
	localPath string
}

func NewUploaderApi(service service.IUploaderService, localPath string) *UploaderApi {
	return &UploaderApi{
		service:   service,
		localPath: localPath,
	}
}

// CheckChunkInfo 步骤1前端请求查看上传数据 查看文件是否存在和之前分片上传的信息
func (m UploaderApi) CheckChunkInfo(c *gin.Context, req dto.ChunkInfoReq) (any, error) {
	//这里也无法验证合法性 因为都是字符数据可造假
	info := m.service.CheckChunkInfo(c, req)
	return gres.Response{
		Data: info,
	}, nil
}

// PutChunk 步骤2分片上传
func (m UploaderApi) PutChunk(c *gin.Context, req dto.ChunkInfoReq) (any, error) {
	//if req.TotalChunks == 1 {
	//	//如果总数只有一个 那么直接保存即可
	//}
	// 尝试从请求中获取分片文件
	if err := m.service.PutChunk(c, "upfile", req); err != nil {
		return nil, err
	}
	//验证文件合法性???实现验证文件合法性文件格式 这里应该无法验证合法性了 因为是二进制数据
	return gres.Response{
		Message: "分片上传成功",
	}, nil
}

// MergeFile 步骤3请求合并文件
func (m UploaderApi) MergeFile(c *gin.Context, req dto.MergeFileReq) (any, error) {
	//合并的时候验证合法性???
	err := m.service.MergeFile(c, req)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "合并成功",
	}, nil
}
