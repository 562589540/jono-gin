package dto

import (
	"encoding/json"
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/internal/constants/enum"
	"strconv"
)

type AttachmentSearchReq struct {
	gdto.PaginateReq        //分页
	FileName         string `json:"fileName" form:"fileName"`           // 文件名,
	FileSize         string `json:"fileSize" form:"fileSize"`           // 文件大小,
	FileType         string `json:"fileType" form:"fileType"`           // 文件类型,
	Path             string `json:"path" form:"path"`                   // 文件地址,
	StorageEngine    string `json:"storageEngine" form:"storageEngine"` // 储存引擎,
	Class            uint   `json:"class" form:"class"`                 // 类别,
	CreatedBy        uint   `json:"createdBy" form:"createdBy"`         // 上传者
}

type AttachmentUpdateReq struct {
	ID int64 `json:"id" binding:"required"`
	AttachmentAddReq
}

type AttachmentAddReq struct {
	FileName      string             `json:"fileName" binding:"required" m:"文件名不能为空"` // 文件名,
	FileSize      string             `json:"fileSize"`                                // 文件大小,
	FileType      string             `json:"fileType"`                                // 文件类型,
	Path          string             `json:"path"`                                    // 文件地址,
	StorageEngine enum.StorageEngine `json:"storageEngine"`                           // 储存引擎,
	Class         uint               `json:"class"`                                   // 类别,
	CreatedBy     uint               `json:"createdBy"`                               // 上传者
}

type Attachment struct {
	ID            int64  `json:"id"`            // ,
	FileName      string `json:"fileName"`      // 文件名,
	FileSize      string `json:"fileSize"`      // 文件大小,
	FileType      string `json:"fileType"`      // 文件类型,
	Path          string `json:"path"`          // 文件地址,
	StorageEngine uint   `json:"storageEngine"` // 储存引擎,
	Class         uint   `json:"class"`         // 类别,
	CreatedBy     uint   `json:"createdBy"`     // 上传者
}

func (a Attachment) MarshalJSON() ([]byte, error) {
	type Alias Attachment
	return json.Marshal(&struct {
		ID string `json:"id"`
		*Alias
	}{
		ID:    strconv.FormatInt(a.ID, 10),
		Alias: (*Alias)(&a),
	})
}
