package mysql

import (
	"GoOriginHttp/util"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/exp/slog"
	"strings"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var MyEngine *xorm.Engine

func SetEngine() {
	user := util.GetVal("mysql", "user")
	ip := util.GetVal("mysql", "ip")
	port := util.GetVal("mysql", "port")
	passwd := util.GetVal("mysql", "passwd")
	database := util.GetVal("mysql", "database")

	uri := strings.Join([]string{ip, port}, ":")
	src := strings.Join([]string{user, ":", passwd, "@tcp(", uri, ")/", database, "?charset=utf8"}, "")
	slog.Info("数据库链接", slog.String("参数", src))
	MyEngine, _ = xorm.NewEngine("mysql", src)

	MyEngine.SetMapper(names.GonicMapper{})
	MyEngine.ShowSQL(true)
	MyEngine.SetTZDatabase(time.Local)

	err := MyEngine.Ping()
	if err != nil {
		slog.Error("创建数据库引擎失败", slog.Any("错误信息", err))
	} else {
		slog.Info("创建数据库引擎成功", slog.Any("MyEngine", MyEngine))
	}

}

func GetSession() *xorm.Session {
	return MyEngine.NewSession()
}
