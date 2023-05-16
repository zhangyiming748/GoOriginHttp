package logic

import (
	"GoOriginHttp/api"
	"GoOriginHttp/model"
	"GoOriginHttp/util"
	"encoding/json"
	"github.com/zhangyiming748/goini"
	"golang.org/x/exp/slog"
	"strconv"
	"time"
)

const (
	configPath = "./conf.ini"
	AMAPURL    = "https://restapi.amap.com/v3/weather/weatherInfo?parameters"
)

type live struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Lives    []struct {
		Province         string `json:"province"`
		City             string `json:"city"`
		Adcode           string `json:"adcode"`
		Weather          string `json:"weather"`
		Temperature      string `json:"temperature"`
		Winddirection    string `json:"winddirection"`
		Windpower        string `json:"windpower"`
		Humidity         string `json:"humidity"`
		Reporttime       string `json:"reporttime"`
		TemperatureFloat string `json:"temperature_float"`
		HumidityFloat    string `json:"humidity_float"`
	} `json:"lives"`
}
type forecast struct {
	Status    string `json:"status"`
	Count     string `json:"count"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Forecasts []struct {
		City       string `json:"city"`
		Adcode     string `json:"adcode"`
		Province   string `json:"province"`
		Reporttime string `json:"reporttime"`
		Casts      []struct {
			Date           string `json:"date"`
			Week           string `json:"week"`
			Dayweather     string `json:"dayweather"`
			Nightweather   string `json:"nightweather"`
			Daytemp        string `json:"daytemp"`
			Nighttemp      string `json:"nighttemp"`
			Daywind        string `json:"daywind"`
			Nightwind      string `json:"nightwind"`
			Daypower       string `json:"daypower"`
			Nightpower     string `json:"nightpower"`
			DaytempFloat   string `json:"daytemp_float"`
			NighttempFloat string `json:"nighttemp_float"`
		} `json:"casts"`
	} `json:"forecasts"`
}

type toAmapWeather struct {
	Key        string `json:"key"`
	City       string `json:"city"`
	Extensions string `json:"extensions"`
	Output     string `json:"output"`
}

func GetWeather(city, extensions string) (res api.CustomerResponse) {
	ext, err := strconv.Atoi(extensions)
	if err != nil {
		return api.CustomerResponse{}
	}
	conf := goini.SetConfig(configPath)
	key, err := conf.GetValue("weather", "key")
	if err != nil {
		slog.Warn("lost key")
	}
	var kind string
	if ext == 1 {
		kind = "base"
	} else if ext == 0 {
		kind = "all"
	} else {
		slog.Warn("参数错误")
	}
	var tam = &toAmapWeather{
		Key:        key,
		City:       City[city],
		Extensions: kind,
		Output:     "JSON",
	}
	b := getFromAmap(tam)
	res.RetData = string(b)
	res.ResStatus = "success"
	res.ResCode = "200"
	return res
}
func getFromAmap(tam *toAmapWeather) []byte {
	var m = map[string]string{
		"key":        tam.Key,
		"city":       tam.City,
		"extensions": tam.Extensions,
		"output":     tam.Output,
	}
	body := util.HttpGetValue(nil, m, AMAPURL)
	if tam.Extensions == "base" {
		var l live
		err := json.Unmarshal(body, &l)
		if err != nil {
			slog.Warn("解析失败")
		}
		localTime, err := time.ParseInLocation("2006-01-02 15:04:05", l.Lives[0].Reporttime, time.Local)
		if err != nil {
			return nil
		}
		//2023-05-05 10:39:20
		var livedb = &model.Live{
			Province:      l.Lives[0].Province,
			City:          l.Lives[0].City,
			Adcode:        l.Lives[0].Adcode,
			Weather:       l.Lives[0].Weather,
			Temperature:   l.Lives[0].Temperature,
			WindDirection: l.Lives[0].Winddirection,
			WindPower:     l.Lives[0].Windpower,
			Humidity:      l.Lives[0].Humidity,
			ReportTime:    localTime,
		}
		livedb.InsertOne()
	} else if tam.Extensions == "all" {
		var fs []model.Forecast
		var all forecast
		err := json.Unmarshal(body, &all)
		if err != nil {
			slog.Warn("解析失败")
		}
		reportTime, _ := time.ParseInLocation("2006-01-02 15:04:05", all.Forecasts[0].Reporttime, time.Local)
		for _, day := range all.Forecasts[0].Casts {
			var f = &model.Forecast{
				Id:           0,
				Province:     all.Forecasts[0].Province,
				City:         all.Forecasts[0].City,
				Adcode:       all.Forecasts[0].Adcode,
				Date:         day.Date,
				Week:         day.Week,
				DayWeather:   day.Dayweather,
				NightWeather: day.Nightweather,
				DayWind:      day.Daywind,
				NightWind:    day.Nightwind,
				DayPower:     day.Daypower,
				NightPower:   day.Nightpower,
				ReportTime:   reportTime,
			}
			fs = append(fs, *f)
		}
		model.InsertForecasts(fs)
		slog.Info("完成一次完整的查询天气请求")
	}
	return body
}
func DeleteAllLive() (res api.CustomerResponse) {
	var one model.Live
	all, err := one.DeleteAll()
	if err != nil {
		slog.Warn("删除出错", slog.Any("错误文本", err))
	}
	res.ResCode = "200"
	res.ResStatus = "success"
	res.RetData = all
	return
}
func GetAllLive() (res api.CustomerResponse) {
	var one model.Live
	if all, err := one.GetAllByXORM(); err != nil {
		res.ResCode = "500"
		res.ResStatus = "查询出错"
		res.RetData = err
	} else {
		res.ResCode = "200"
		res.ResStatus = "查询成功"
		res.RetData = all
	}
	return res
}
