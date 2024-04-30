package gclient

import (
	"bytes"
	"io"
	"net/http"
	"sync"
	"time"
)

type HttpClient struct {
	Client *http.Client
}

var (
	instance *HttpClient
	once     sync.Once
)

// NewHttpClient 创建并返回一个配置好的 HTTP 客户端实例。
func NewHttpClient() *HttpClient {
	once.Do(func() {
		instance = &HttpClient{
			Client: &http.Client{
				Timeout: time.Second * 10, // 设置请求超时时间
			},
		}
	})
	return instance
}

// request 是一个通用的请求方法，支持任意 HTTP 方法，请求体和自定义头部。
func (c *HttpClient) request(method, url string, body io.Reader, headers map[string]string) (string, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//读取响应体
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

// Get 发送一个 GET 请求。
func (c *HttpClient) Get(url string, headers map[string]string) (string, error) {
	return c.request("GET", url, nil, headers)
}

// Post 发送一个 POST 请求，内容类型默认为 application/json。
func (c *HttpClient) Post(url string, data []byte, headers map[string]string) (string, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json" // 默认设置为 JSON，可以通过 headers 参数覆盖
	return c.request("POST", url, bytes.NewBuffer(data), headers)
}

// Put 发送一个 PUT 请求。
func (c *HttpClient) Put(url string, data []byte, headers map[string]string) (string, error) {
	return c.request("PUT", url, bytes.NewBuffer(data), headers)
}

// Delete 发送一个 DELETE 请求。
func (c *HttpClient) Delete(url string, headers map[string]string) (string, error) {
	return c.request("DELETE", url, nil, headers)
}

//    // 发送 GET 请求
//    response, err := gclient.Get("https://api.example.com/data", nil)
//    if err != nil {
//        fmt.Println("GET error:", err)
//    } else {
//        fmt.Println("GET response:", response)
//    }
//
//    // 发送 POST 请求
//    jsonData := []byte(`{"name":"John", "age":30}`)
//    postResponse, err := gclient.Post("https://api.example.com/data", jsonData, nil)
//    if err != nil {
//        fmt.Println("POST error:", err)
//    } else {
//        fmt.Println("POST response:", postResponse)
//    }
//
//    // 发送 PUT 请求
//    putData := []byte(`{"name":"John Updated", "age":31}`)
//    putResponse, err := gclient.Put("https://api.example.com/data/1", putData, nil)
//    if err != nil {
//        fmt.Println("PUT error:", err)
//    } else {
//        fmt.Println("PUT response:", putResponse)
//    }
//
//    // 发送 DELETE 请求
//    deleteResponse, err := gclient.Delete("https://api.example.com/data/1", nil)
//    if err != nil {
//        fmt.Println("DELETE error:", err)
//    } else {
//        fmt.Println("DELETE response:", deleteResponse)
//    }
