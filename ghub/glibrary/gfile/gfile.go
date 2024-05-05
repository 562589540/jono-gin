package gfile

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub/gutils"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// WriteToFile 写入文件
// 如果目的是追加内容，应移除 os.O_TRUNC 标志。
func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	n, _ := f.Seek(0, os.SEEK_END)
	_, err = f.WriteAt([]byte(content), n)
	defer f.Close()
	return err
}

// FileIsExisted 文件或文件夹是否存在
func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// IsNotExistMkDir 检查文件夹是否存在  如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := !FileIsExisted(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// GetType 获取文件类型
func GetType(p string) (result string, err error) {
	file, err := os.Open(p)
	if err != nil {
		gutils.Error(err)
		return
	}
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		gutils.Error(err)
		return
	}
	filetype := http.DetectContentType(buff)
	return filetype, nil
}

// CheckAndCreateFile
// 检查文件是否存在不存在创建并且出数数据
// file.WriteString("initial content\n")
func CheckAndCreateFile(filePath string, fn func(file *os.File) error) error {
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 文件不存在，创建文件以及目录
		if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		file, fErr := os.Create(filePath)
		if fErr != nil {
			return fErr
		}
		defer file.Close()
		return fn(file)
	}
	return nil
}

// CheckAndDeleteFile 检查文件是否存在，如果存在则删除
func CheckAndDeleteFile(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		return os.Remove(filePath)
	} else if os.IsNotExist(err) {
		return nil
	} else {
		// 其他错误
		return err
	}
}

// CheckAndClearFileIfExist 检查文件是否存在，如果存在则清空文件内容
func CheckAndClearFileIfExist(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		return os.Truncate(filePath, 0)
	} else if os.IsNotExist(err) {
		return nil
	} else {
		// 其他错误
		return err
	}
}

// CreateFileAndWrite 创建并且写入
func CreateFileAndWrite(filePath string, fn func(file *os.File) error) error {
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	file, fErr := os.Create(filePath)
	if fErr != nil {
		return fErr
	}
	defer file.Close()
	return fn(file)
}

func FormatBytes(bytes int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
	)

	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", bytes) // Byte is the smallest unit
	}
}
