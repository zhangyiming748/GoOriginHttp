package model

import (
	"GoOriginHttp/mysql"
	_ "GoOriginHttp/util"
	"fmt"
	"golang.org/x/exp/slog"
	"testing"
)

func TestSelectLive(t *testing.T) {

	mysql.SetEngine()
	var l Live
	count, err := l.Count()
	if err != nil {
		return
	}
	slog.Info("count", slog.Int64("count", count))
	ls, _ := l.GetAll()
	for _, v := range ls {
		adcode := string(v["adcode"])
		fmt.Println(adcode)
		//zpretty.Println(v)
	}
}
