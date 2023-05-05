package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangyiming748/goini"
	"golang.org/x/exp/slog"
	"strings"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

const (
	configurePath = "./conf.ini"
)

var MyEngine *xorm.Engine

func SetEngine() {
	conf := goini.SetConfig(configurePath)
	user, err := conf.GetValue("mysql", "user")
	ip, err := conf.GetValue("mysql", "ip")
	port, err := conf.GetValue("mysql", "port")
	passwd, err := conf.GetValue("mysql", "passwd")
	database, err := conf.GetValue("mysql", "database")
	if err != nil {
		slog.Warn("没有找到配置")
	}
	uri := strings.Join([]string{ip, port}, ":")
	src := strings.Join([]string{user, ":", passwd, "@tcp(", uri, ")/", database, "?charset=utf8"}, "")
	slog.Info("数据库链接", slog.String("参数", src))
	if MyEngine, err = xorm.NewEngine("mysql", src); err != nil {
		slog.Error("创建数据库引擎失败")
		panic(err)
	}
	MyEngine.SetMapper(names.GonicMapper{})
	MyEngine.ShowSQL(true)
	MyEngine.SetTZDatabase(time.Local)
	//err = MyEngine.Sync2(new(model.Live), new(model.Forecast))
	//if err != nil {
	//	slog.Error("同步数据表出错", slog.Any("错误原文", err))
	//	return
	//}
	err = MyEngine.Ping()
	if err != nil {
		slog.Error("创建数据库引擎失败", slog.Any("错误信息", err))
	} else {
		slog.Info("创建数据库引擎成功", slog.Any("MyEngine", MyEngine))
	}

}

func GetSession() *xorm.Session {
	return MyEngine.NewSession()
}
