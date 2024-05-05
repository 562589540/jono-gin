package dto

type ChunkInfoReq struct {
	ChunkNumber      int    `json:"chunkNumber" form:"chunkNumber"`                                                  //总块数量
	ChunkSize        int64  `json:"chunkSize" form:"chunkSize"`                                                      //块大小
	CurrentChunkSize int    `json:"currentChunkSize" form:"currentChunkSize"`                                        //当前块大小
	TotalSize        int    `json:"totalSize" form:"totalSize"`                                                      //总大小
	Identifier       string `json:"identifier" form:"identifier" binding:"required,md5" m:"哈希值不能为空" md5_m:"哈希值格式错误"` //md5
	Filename         string `json:"filename" form:"filename"`                                                        //文件名称
	RelativePath     string `json:"relativePath" form:"relativePath"`                                                //相对路径
	TotalChunks      int    `json:"totalChunks" form:"totalChunks"`                                                  //总块
}

type ChunkInfoRes struct {
	SkipUpload     bool  `json:"skipUpload"`     //文件已经存在
	UploadedChunks []int `json:"uploadedChunks"` //文件块列表
}

type MergeFileReq struct {
	UniqueIdentifier string `json:"uniqueIdentifier" binding:"required,md5" m:"哈希值不能为空" md5_m:"哈希值格式错误"` //哈希值
	FileType         string `json:"fileType"`                                                            //文件类型 后端也要验证
	Name             string `json:"name"`                                                                //文件名称
	RelativePath     string `json:"relativePath"`                                                        //文件名称
	Size             int    `json:"size"`                                                                //文件大小
}
