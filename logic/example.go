package logic

import (
	"GoOriginHttp/api"
	"golang.org/x/exp/slog"
	"time"
)

var chooseProb []float64

/*
*

	首页推荐
*/
type Persion struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func FindByName(p *Persion) (res api.CustomerResponse) {
	res.ResCode = api.APPRESPONSE_CODE_SUCCESS
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.RESPONSE_CODE_FAIL
			res.ResStatus = api.REQUEST_ERR
			slog.Error("FindByName", slog.Any("错误", err))
		}
	}()

	var persion Persion
	//最多等待400ms
	c := make(chan Persion)
	go func() {
		c <- persion
	}()
	select {
	case persion = <-c:
	case <-time.After(400 * time.Millisecond):
		slog.Error("400ms time out!")
	}
	res.RetData = res
	slog.Debug("根据姓名查找", slog.Any("in", *p), slog.Any("out", res))
	return
}
