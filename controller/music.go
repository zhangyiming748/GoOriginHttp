package controller

import (
	"GoOriginHttp/api"
	"GoOriginHttp/logic"
	"net/http"
)

func GetTopMusic(r *http.Request, w http.ResponseWriter) (res api.CustomerResponse) {
	query := r.URL.Query()
	keyword := query.Get("keyword")
	res = logic.GetTopByKeyword(keyword)
	return res
}
