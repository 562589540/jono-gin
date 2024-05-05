package gutils

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/jinzhu/copier"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

type Utils struct {
}

// AppendError 错误链
func AppendError(existErr, newErr error) error {
	if existErr == nil {
		return newErr
	}

	return fmt.Errorf("%v,%w", existErr, newErr)
}

// ErrorPanic 报错并且停止运行
func ErrorPanic(err error) {
	if ghub.Log != nil {
		ghub.Log.Error(err.Error())
	}
	panic(err.Error())
}

// GetLocalIP 服务端ip
func GetLocalIP() (ip string, err error) {
	var addrs []net.Addr
	addrs, err = net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// Info log
func Info(log string) {
	if ghub.Log != nil {
		ghub.Log.Info(log)
	} else {
		fmt.Println(log)
	}
}

func Error(err error) {
	if ghub.Log != nil {
		ghub.Log.Error(err)
	} else {
		fmt.Println(err)
	}
}

func CheckError(err error) {
	if err != nil {
		if ghub.Log != nil {
			ghub.Log.Error(err)
		} else {
			fmt.Println(err)
		}
	}
}

// GetCityByIp 根据 IP 地址获取所属城市名称。
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "::1" || ip == "127.0.0.1" {
		return "内网IP"
	}

	url := fmt.Sprintf("http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=%s", ip)
	response, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	// 转换字符集从 GBK 到 UTF-8
	utf8Reader := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
	utf8Response, err := io.ReadAll(utf8Reader)
	if err != nil {
		return ""
	}

	var result map[string]interface{}
	err = json.Unmarshal(utf8Response, &result)
	if err != nil {
		return ""
	}

	province, okPro := result["pro"].(string)
	city, okCity := result["city"].(string)
	if okPro && okCity {
		return fmt.Sprintf("%s %s", province, city)
	}

	return ""
}

// Copy 复制结构体
func Copy(dst any, src any) error {
	if err := copier.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

// ParseTimeInterval 获取时间数组结构
func ParseTimeInterval(c *gin.Context, timeStr string) (mTimes []time.Time, err error) {
	startTimeStr := c.Query(timeStr + "[0]")
	endTimeStr := c.Query(timeStr + "[1]")

	if startTimeStr != "" {
		startTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			return nil, err
		}
		mTimes = append(mTimes, startTime)
	}

	if endTimeStr != "" {
		endTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			return nil, err
		}
		mTimes = append(mTimes, endTime)
	}

	return mTimes, nil
}

// Contains 检查切片中是否包含某个元素
func Contains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// FormatDuration 时间输出显示
func FormatDuration(d time.Duration) string {
	var result strings.Builder

	totalSeconds := int(d.Seconds())
	seconds := totalSeconds % 60
	minutes := (totalSeconds / 60) % 60
	hours := (totalSeconds / 3600) % 24
	days := totalSeconds / 86400

	if days > 0 {
		fmt.Fprintf(&result, "%d天 ", days)
	}
	if days > 0 || hours > 0 {
		fmt.Fprintf(&result, "%d时 ", hours)
	}
	if days > 0 || hours > 0 || minutes > 0 {
		fmt.Fprintf(&result, "%d分 ", minutes)
	}
	if days > 0 || hours > 0 || minutes > 0 || seconds > 0 {
		fmt.Fprintf(&result, "%d秒", seconds)
	}

	return result.String()
}

// GetFriendlyOSName 系统转换
func GetFriendlyOSName(os string) string {
	switch os {
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	case "freebsd":
		return "FreeBSD"
	case "openbsd":
		return "OpenBSD"
	case "netbsd":
		return "NetBSD"
	default:
		return os // 如果是未知系统，返回原始值
	}
}

// BytesToGB bytesToGB 转换字节为千兆字节
func BytesToGB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}
