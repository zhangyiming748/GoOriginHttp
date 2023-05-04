package controller

import (
	"net/http"
	"recommend/api"
)

/*
Get 方法根据姓名查询全部信息
*/
func GetWeather(r *http.Request, w http.ResponseWriter) (res api.CustomerResponse) {
	query := r.URL.Query()
	city := query.Get("City")
	// 获取URL路径参数
	///user/10?name=john&age=30
	//like user/:id
	//vars := mux.Vars(r)
	//id := vars["id"]
	
	res.RetData = city
	return
}
