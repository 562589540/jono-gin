package gutils

import (
	"encoding/base64"
	"fmt"
	"github.com/562589540/jono-gin/ghub/gbootstrap"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SaveBase64Image 保存从 base64 编码字符串解码的图像到指定目录下的按日期分类的子目录，返回文件路径和可能出现的错误。
func SaveBase64Image(b64data string, baseUploadDir string) (string, error) {
	if baseUploadDir == "" {
		//使用默认路径
		baseUploadDir = filepath.Join(gbootstrap.Cfg.Path.ResourcePath, gbootstrap.Cfg.Path.UploadsPath, gbootstrap.Cfg.Path.AvatarPath)
	}
	base64Data := strings.Split(b64data, ",")[1]
	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}

	// 根据当前日期创建文件夹路径，格式如 2023-01-25
	dateDir := time.Now().Format("2006-01-02")
	uploadDir := filepath.Join(baseUploadDir, dateDir)

	// 确保目录存在
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, 0755)
		if err != nil {
			return "", err
		}
	}

	// 使用 UUID 生成唯一的文件名
	uniqueID := uuid.New()
	filename := filepath.Join(uploadDir, fmt.Sprintf("avatar-%s.png", uniqueID))

	// 写入文件
	err = os.WriteFile(filename, decoded, 0666)
	if err != nil {
		return "", err
	}

	// 计算相对路径
	relPath, err := filepath.Rel(gbootstrap.Cfg.Path.ResourcePath, filename)
	if err != nil {
		return "", err
	}
	// 替换路径分隔符
	relPath = filepath.ToSlash(relPath)
	return filepath.Join(gbootstrap.Cfg.Path.Static, relPath), nil
}
