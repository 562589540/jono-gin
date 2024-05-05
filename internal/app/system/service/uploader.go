package service

import (
	"context"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/gin-gonic/gin"
)

type IUploaderService interface {
	CheckChunkInfo(ctx context.Context, data dto.ChunkInfoReq) *dto.ChunkInfoRes
	MergeFile(ctx context.Context, data dto.MergeFileReq) error
	PutChunk(c *gin.Context, formFile string, req dto.ChunkInfoReq) error
}
