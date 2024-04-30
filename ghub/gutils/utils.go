package gutils

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/goccy/go-json"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"net"
	"net/http"
)

type Utils struct {
}

func AppendError(existErr, newErr error) error {
	if existErr == nil {
		return newErr
	}

	return fmt.Errorf("%v,%w", existErr, newErr)
}

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

func Error(err error) {
	if ghub.Log != nil {
		ghub.Log.Error(err.Error())
	} else {
		fmt.Println(err)
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
