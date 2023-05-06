package controller

import (
	"GoOriginHttp/api"
	"GoOriginHttp/logic"
	"net/http"
)

/*
Get 方法查询天气
curl --location --request GET 'http://127.0.0.1:9090/api/v1/GetWeathe?City=<City>' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'
*/
func GetWeather(r *http.Request, w http.ResponseWriter) (res api.CustomerResponse) {
	query := r.URL.Query()
	city := query.Get("City") //城市的中文名
	extensions := query.Get("extensions")
	// 获取URL路径参数
	///user/10?name=john&age=30
	//like user/:id
	//vars := mux.Vars(r)
	//id := vars["id"]
	res = logic.GetWeather(city, extensions)
	return
}

/*
删除全部实时天气记录
curl --location --request GET 'http://127.0.0.1:9090/api/v1/DeleteAllLive' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'
*/
func DeleteAllLive(r *http.Request, w http.ResponseWriter) (res api.CustomerResponse) {
	res = logic.DeleteAllLive()
	return
}

/*
获取城市代码列表
curl --location --request GET 'http://127.0.0.1:9090/api/v1/GetCity' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'
*/
func GetCity(r *http.Request, w http.ResponseWriter) (res api.CustomerResponse) {
	res.ResCode = "200"
	res.ResStatus = "成功"
	res.RetData = logic.City
	return
}
