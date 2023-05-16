package logic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const BaiduTTS = "https://xiaoapi.cn/API/tts.php"

type BaiduTTSREQ struct {
	Code int `json:"code"`
	Data struct {
		Msg string `json:"msg"`
		Mp3 string `json:"mp3"`
	} `json:"data"`
	Tips string `json:"tips"`
}

func TTS(msg string) {
	voice := getTTS(msg)
	js := fmt.Sprintf(voice)
	var bt BaiduTTSREQ
	err := json.Unmarshal([]byte(js), &bt)
	if err != nil {
		fmt.Println("fail")
	}
	fmt.Printf("%+v\n", bt)
}
func getTTS(msg string) string {
	msg = strings.Join([]string{"<", msg, ">"}, "")
	text := url.QueryEscape(msg)
	uri := strings.Join([]string{BaiduTTS, "?msg=", text}, "")
	fmt.Println(uri)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, uri, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	s := string(body)
	return s
}
