package model

import (
	"GoOriginHttp/mysql"
	"golang.org/x/exp/slog"
	"time"
)

type Forecast struct {
	Id           int       `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	Province     string    `xorm:"comment('省份名') VARCHAR(255)" json:"province"`
	City         string    `xorm:"comment('城市名') VARCHAR(255)" json:"city"`
	Adcode       string    `xorm:"comment('区域编码') VARCHAR(255)" json:"adcode"`
	Date         string    `xorm:"comment('预报日期') VARCHAR(255)" json:"date"`
	Week         string    `xorm:"comment('星期') VARCHAR(255)" json:"week"`
	DayWeather   string    `xorm:"comment('白天天气') VARCHAR(255)" json:"dayWeather"`
	NightWeather string    `xorm:"comment('夜间天气') VARCHAR(255)" json:"nightWeather"`
	DayWind      string    `xorm:"comment('白天风向') VARCHAR(255)" json:"dayWind"`
	NightWind    string    `xorm:"comment('夜间风向') VARCHAR(255)" json:"nightWind"`
	DayPower     string    `xorm:"comment('白天风力') VARCHAR(255)" json:"dayPower"`
	NightPower   string    `xorm:"comment('夜间风力') VARCHAR(255)" json:"nightPower"`
	ReportTime   time.Time `xorm:"comment('预报发布时间') DateTime" json:"reportTime"`
	UpdateTime   time.Time `xorm:"updated comment('更新时间) DateTime" json:"update_time"`
	CreateTime   time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime   time.Time `xorm:"deleted comment('创建时间') DateTime" json:"delete_time"`
}

func (f Forecast) FindByCityName(city string) []Forecast {
	var forecasts []Forecast
	err := mysql.GetSession().Where("name = ?", city).Find(&forecasts)
	if err != nil {
		return []Forecast{}
	}
	return forecasts
}
func InsertForecasts(fs []Forecast) {
	insert, err := mysql.GetSession().Insert(fs)
	if err != nil {
		return
	} else {
		slog.Info("插入数据", slog.Int64("条数", insert))
	}
}
