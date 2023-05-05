package util

import (
	"bytes"
	"encoding/json"
	"golang.org/x/exp/slog"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

func httpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		slog.Warn("httpGet", slog.Any("错误", err))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		slog.Warn("httpGet读返回体", slog.Any("错误", err), slog.Any("返回体", body))
	}
	result := string(body)
	slog.Info("http get done", slog.String("返回内容", result))
	return body
}
func HttpGetValue(addHeaders map[string]string, data map[string]string, urlPath string) []byte {
	params := url.Values{}
	urlInfo, err := url.Parse(urlPath)
	if err != nil {
		slog.Warn("解析url出错", slog.Any("错误内容", err))
	}
	for dataKey, dataVal := range data {
		params.Set(dataKey, dataVal)
	}
	urlInfo.RawQuery = params.Encode()
	fullUrl := urlInfo.String()
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		slog.Warn("包装url出错", slog.Any("错误内容", err))
	}
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Warn("发送http请求出错", slog.Any("错误内容", err))
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Warn("读返回体出错", slog.Any("错误内容", err), slog.String("返回体", string(body)))
	}
	slog.Info("http get done", slog.String("返回体", string(body)))
	return body
}

func httpPost(url string, fields ...string) []byte {
	var s string = ""
	for i, v := range fields {
		val := ""
		if i%2 == 0 {
			val = v + "="
		} else {
			if i != len(fields)-1 {
				val = v + "&"
			} else {
				val = v
			}
		}
		s = s + val
	}
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(s))
	if err != nil {
		slog.Warn("http post 请求出错", slog.Any("错误内容", err))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Warn("读返回体出错", slog.Any("错误内容", err))
	}
	return body
}
func HttpPostJson(addHeaders map[string]string, data interface{}, urlPath string) []byte {
	bytesData, err := json.Marshal(data)
	if err != nil {
		slog.Warn("编码json出错", slog.Any("错误内容", err))
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", urlPath, reader)
	if err != nil {
		slog.Warn("包装post请求出错", slog.Any("错误内容", err))
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Warn("发送post请求出错", slog.Any("错误内容", err))
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Warn("读返回体出错", slog.Any("错误内容", err), slog.String("返回体", string(body)))
	}
	slog.Info("http post json done", slog.String("返回体", string(body)))
	return body
}
func HttpProxyFileUpload(file *multipart.FileHeader, fileKey string, addFields map[string]string, addHeaders map[string]string, urlPath string) []byte {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	formFile, err := writer.CreateFormFile(fileKey, file.Filename)
	if err != nil {
		slog.Warn("编码json出错", slog.Any("错误内容", err))
	}

	// 从文件读取数据，写入表单
	srcFile, err := file.Open()
	if err != nil {
		slog.Warn("编码json出错", slog.Any("错误内容", err))
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		slog.Warn("编码json出错", slog.Any("错误内容", err))
	}
	for fieldKey, fieldVal := range addFields {
		if err = writer.WriteField(fieldKey, fieldVal); err != nil {
			slog.Warn("编码json出错", slog.Any("错误内容", err))
		}
	}
	// 发送表单
	contentType := writer.FormDataContentType()
	writer.Close() // 发送之前必须调用Close()以写入结尾行
	req, err := http.NewRequest("POST", urlPath, buf)
	if err != nil {
		slog.Warn("编码json出错", slog.Any("错误内容", err))
	}
	req.Header.Set("Content-Type", contentType)
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		slog.Warn("编码json出错", slog.Any("错误内容", err))
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, _ := io.ReadAll(resp.Body)
	return body
}
