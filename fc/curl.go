package fc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// curlGet 模拟 PHP curl GET 请求
func CurlGet(url string) (string, error) {

	// 使用示例
	options := CurlOptions{
		URL:     url,
		Timeout: 15 * time.Second,
		Headers: map[string]string{
			"User-Agent":   "MyGoClient/1.0",
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
		// Proxy: "http://proxy.example.com:8080",
	}

	body, code, err := CurlGetUp(options)
	if err != nil {
		err = fmt.Errorf("statusCode:%d , 请求失败: %s", code, err.Error())
		return "", err
	}

	return body, nil

}

// CurlOptions 类似 PHP curl 的选项配置
type CurlOptions struct {
	URL     string
	Timeout time.Duration
	Headers map[string]string
	Proxy   string
}

// CurlGet 增强版 GET 请求
func CurlGetUp(options CurlOptions) (string, int, error) {
	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: options.Timeout,
	}

	// 设置代理（如果需要）
	if options.Proxy != "" {
		// 将字符串代理地址转换为 *url.URL
		proxyURL, err := url.Parse(options.Proxy)
		if err != nil {
			return "", 0, fmt.Errorf("代理地址解析失败: %v", err)
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		client.Transport = transport
	}

	// 创建请求
	req, err := http.NewRequest("GET", options.URL, nil)
	if err != nil {
		return "", 0, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, fmt.Errorf("读取响应失败: %v", err)
	}

	return string(body), resp.StatusCode, nil
}

// CurlPostOptions POST 请求配置
type CurlPostOptions struct {
	URL         string
	Data        interface{}       // POST 数据
	Headers     map[string]string // 请求头
	Files       map[string]string // 文件上传 {字段名: 文件路径}
	Timeout     time.Duration
	ContentType string // 可选: form, json, multipart
}

// CurlPost 模拟 PHP cURL POST 请求
func CurlPostUp(options CurlPostOptions) (string, int, error) {
	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: options.Timeout,
	}

	var reqBody io.Reader
	var contentType string

	// 根据 ContentType 准备请求体
	switch strings.ToLower(options.ContentType) {
	case "json":
		// JSON 数据
		jsonData, err := json.Marshal(options.Data)
		if err != nil {
			return "", 0, fmt.Errorf("JSON编码失败: %v", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
		contentType = "application/json"

	case "multipart":
		// 表单数据 + 文件上传
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// 添加普通字段
		if data, ok := options.Data.(map[string]string); ok {
			for key, value := range data {
				writer.WriteField(key, value)
			}
		}

		// 添加文件
		for fieldName, filePath := range options.Files {
			file, err := os.Open(filePath)
			if err != nil {
				return "", 0, fmt.Errorf("打开文件失败: %v", err)
			}
			defer file.Close()

			part, err := writer.CreateFormFile(fieldName, filepath.Base(filePath))
			if err != nil {
				return "", 0, fmt.Errorf("创建表单文件失败: %v", err)
			}

			_, err = io.Copy(part, file)
			if err != nil {
				return "", 0, fmt.Errorf("复制文件内容失败: %v", err)
			}
		}

		writer.Close()
		reqBody = body
		contentType = writer.FormDataContentType()

	default:
		// 默认表单数据
		if data, ok := options.Data.(map[string]string); ok {
			formData := url.Values{}
			for key, value := range data {
				formData.Set(key, value)
			}
			reqBody = strings.NewReader(formData.Encode())
			contentType = "application/x-www-form-urlencoded"
		} else if data, ok := options.Data.(string); ok {
			// 直接字符串数据
			reqBody = strings.NewReader(data)
			contentType = "application/x-www-form-urlencoded"
		}
	}

	// 创建请求
	req, err := http.NewRequest("POST", options.URL, reqBody)
	if err != nil {
		return "", 0, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置 Content-Type
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	// 设置其他请求头
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("send request failed: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, fmt.Errorf("read response failed: %v", err)
	}

	return string(body), resp.StatusCode, nil
}

// CurlPost 适用于简单的 POST 请求，数据以JSON 格式发送
func CurlPost(url string, data interface{}) (string, error) {
	options1 := CurlPostOptions{
		URL:  url,
		Data: data,
		Headers: map[string]string{
			"User-Agent": "Go-Curl-Post/1.0",
			"Accept":     "application/json",
		},
		Timeout: 30 * time.Second,
	}

	response1, status1, err1 := CurlPost(options1)
	if err1 != nil {
		return "", fmt.Errorf("statusCode:%d , err: %s", status1, err1.Error())
	}
	return response1, nil
}

// CurlPostParams 适用于简单的 POST 请求，数据以 表单格式发送
func CurlPostParams(url string, data map[string]interface{}) (string, error) {

	options2 := CurlPostOptions{
		URL:         url,
		Data:        data,
		ContentType: "json",
		Headers:     map[string]string{},
		Timeout:     30 * time.Second,
	}

	response2, status2, err2 := CurlPost(options2)
	if err2 != nil {
		return "", fmt.Errorf("statusCode:%d , err:%s", status2, err2.Error())
	}
	return response2, nil
}
