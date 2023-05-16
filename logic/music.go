package logic

import (
	"GoOriginHttp/api"
	"GoOriginHttp/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	URL = "https://xiaoapi.cn/API/yy_sq.php"
)

type Music struct {
	Code    int    `json:"code"`
	Cover   string `json:"cover"`
	Name    string `json:"name"`
	Singer  string `json:"singer"`
	Quality string `json:"quality"`
	Url     string `json:"url"`
	Tips    string `json:"tips"`
}
type MusicList struct {
	Code int `json:"code"`
	List []struct {
		Name   string `json:"name"`
		Singer string `json:"singer"`
	} `json:"list"`
	Msg  string `json:"msg"`
	Tips string `json:"tips"`
}

// 直接返回搜索中排名第一的歌曲
func GetTopByKeyword(keyword string) (res api.CustomerResponse) {
	msg := url.QueryEscape(keyword)
	req, err := GetOne(msg)
	if err != nil {
		return api.CustomerResponse{}
	}
	sreq := string(req)
	//fmt.Println(sreq)
	var m Music
	json.Unmarshal([]byte(sreq), &m)
	code := strconv.Itoa(m.Code)
	fmt.Printf("code is %v\n", m.Code)
	fmt.Printf("cover is %v\n", m.Cover)
	fmt.Printf("singer is %v\n", m.Singer)
	fmt.Printf("quality is %v\n", m.Quality)
	fmt.Printf("url is %v\n", m.Url)
	fmt.Printf("tips is %v\n", m.Tips)
	one := model.Music{
		Id:      0,
		Code:    code,
		Cover:   m.Cover,
		Name:    m.Name,
		Singer:  m.Singer,
		Quality: m.Quality,
		Url:     m.Url,
		Tips:    m.Tips,
	}
	_, err = one.InsertOne()
	if err != nil {
		res.ResCode = "500"
		res.ResStatus = "失败"
		res.RetData = err
		return res
	}
	res.ResCode = "200"
	res.ResStatus = "成功"
	res.RetData = m
	return res
}
func GetOne(keyword string) ([]byte, error) {

	uri := strings.Join([]string{URL, "?msg=", keyword, "&type=json&n=1"}, "")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "xiaoapi.cn")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, err
}

// 按照关键词搜索歌曲
func SearchByKeyword(keyword string) {

}

// 显示指定索引的歌曲
func SelectByIndex(keyword, index string) {

}
