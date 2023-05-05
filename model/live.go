package model

import (
	"GoOriginHttp/mysql"
	"golang.org/x/exp/slog"
	"time"
)

type Live struct {
	Id            int       `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	Province      string    `xorm:"comment('省份名') VARCHAR(255)" json:"province"`
	City          string    `xorm:"comment('城市名') VARCHAR(255)" json:"city"`
	Adcode        string    `xorm:"comment('区域编码') VARCHAR(255)" json:"adcode"`
	Weather       string    `xorm:"comment('天气现象') VARCHAR(255)" json:"weather"`
	Temperature   string    `xorm:"comment('实时温度') VARCHAR(255)" json:"temperature"`
	WindDirection string    `xorm:"comment('风向描述') VARCHAR(255)" json:"windDirection"`
	WindPower     string    `xorm:"comment('风力级别') VARCHAR(255)" json:"windPower"`
	Humidity      string    `xorm:"comment('空气湿度') VARCHAR(255)" json:"humidity"`
	ReportTime    time.Time `xorm:"comment('数据发布时间') DateTime" json:"reportTime"`
	UpdateTime    time.Time `xorm:"updated comment('更新时间) DateTime" json:"update_time"`
	CreateTime    time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime    time.Time `xorm:"deleted comment('创建时间') DateTime" json:"delete_time"`
}

func (l Live) FindByCityName(city string) []Live {
	var lives []Live
	err := mysql.GetSession().Where("name = ?", city).Find(&lives)
	if err != nil {
		return []Live{}
	}
	return lives
}
func (l Live) InsertOne() {
	insert, err := mysql.GetSession().Insert(l)
	if err != nil {
		return
	} else {
		slog.Info("插入数据", slog.Int64("条数", insert))
	}
}
func (l Live) DeleteAll() ([]map[string][]byte, error) {
	del := "TRUNCATE TABLE live;"
	slog.Info("执行删库跑路命令")
	return mysql.GetSession().Query(del)
}
