package model

import (
	"time"
)

type Chunk struct {
	ID        int64  `gorm:"primaryKey"`                    // 使用雪花ID作为主键
	Number    int    `gorm:"comment:分片序号"`                  // 分片序号
	Path      string `gorm:"comment:分片存储路径"`                // 存储路径
	Size      int64  `gorm:"comment:分片大小"`                  // 分片大小
	Hash      string `gorm:"size:255;index;comment:分片的哈希值"` // 哈希值，用于校验完整性
	Status    int8   `gorm:"comment:分片的上传状态"`               // 如 pending, uploaded, error
	CreatedAt time.Time
}
