package controller

import (
	"GoOriginHttp/api"
	"GoOriginHttp/logic"
	"encoding/json"
	"io"
	"net/http"
)

/*
Get 方法根据姓名查询全部信息
*/
func GetPersionInfo(r *http.Request, w http.ResponseWriter) (res api.CustomerResponse) {
	query := r.URL.Query()
	name := query.Get("name")
	// 获取URL路径参数
	///user/10?name=john&age=30
	//like user/:id
	//vars := mux.Vars(r)
	//id := vars["id"]
	if name == "zen" {
		res.RetData = "done!"
		return
	}
	var p = &logic.Persion{
		Name: name,
	}
	res = logic.FindByName(p)
	return
}

/*
POST方法根据姓名查询全部信息
*/
func PostPersionInfo(r *http.Request, w http.ResponseWriter) (res api.CustomerResponse) {
	defer func() {
		res.RetData = "2"
	}()

	// 设置MIME类型
	w.Header().Set("Content-Type", "application/json")
	// 读取请求体
	body, _ := io.ReadAll(r.Body)

	// 解析JSON
	var req logic.Persion
	err := json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	res = logic.FindByName(&req)
	return
}
