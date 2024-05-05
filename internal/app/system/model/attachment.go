package model

import (
	"github.com/562589540/jono-gin/internal/constants/enum"
	"time"
)

type Attachment struct {
	ID            int64              `gorm:"primarykey"`
	FileName      string             `gorm:"size:50;not null;comment:文件名"`
	FileSize      string             `gorm:"size:50;comment:文件大小"`
	FileType      string             `gorm:"size:50;comment:文件类型"`
	Path          string             `gorm:"size:255;comment:文件地址"`
	StorageEngine enum.StorageEngine `gorm:"comment:储存引擎"`                          //local cos oss 枚举类型
	Hash          string             `gorm:"size:255;not null;index;comment:文件哈希值"` // 哈希值，用于校验完整性
	Class         uint               `gorm:"comment:类别"`
	CreatedBy     uint               `gorm:"comment:上传者"`
	Status        int8               `gorm:"comment:文件上传状态"` // 如 pending, completed, error
	TotalChunks   int                `gorm:"comment:总分片大小"`
	CreatedAt     time.Time
}
