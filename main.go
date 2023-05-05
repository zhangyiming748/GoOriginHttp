package main

import (
	"GoOriginHttp/api"
	"GoOriginHttp/controller"
	"GoOriginHttp/model"
	"GoOriginHttp/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"github.com/zhangyiming748/goini"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"os"

	"strings"
	"time"
)

var (
	//url_prefix = "/api/recommend"
	url_prefix = "/api"
)
var logLevel = map[string]slog.Level{
	"Debug": slog.LevelDebug,
	"Info":  slog.LevelInfo,
	"Warn":  slog.LevelWarn,
	"Error": slog.LevelError,
}

func SetLog(level string) {
	var opt = slog.HandlerOptions{ // 自定义option
		AddSource: true,
		Level:     logLevel[level], // slog 默认日志级别是 info
	}
	file := "normal.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	mylog := slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
	slog.SetDefault(mylog)
}

func main() {
	conf := goini.SetConfig("./conf.ini")
	if level, err := conf.GetValue("log", "level"); err != nil {
		SetLog("Debug")
	} else {
		SetLog(level)
	}
	router := makeRouters()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	mysql.SetEngine()
	model.SyncLive()
	model.SyncForecast()
	n := negroni.New(negroni.NewRecovery())
	n.Use(c)
	n.UseHandler(router)
	s := &http.Server{
		Addr:           ":9090",
		Handler:        n,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	slog.Info("服务启动成功", slog.String("端口号", s.Addr))
	if err := s.ListenAndServe(); err != nil {
		slog.Warn("fall")
	} else {
		slog.Info("success")
	}

}

func makeRouters() *mux.Router {

	jsonrender := render.New(render.Options{UnEscapeHTML: false})

	wrapper := func(apphandler func(*http.Request, http.ResponseWriter) api.CustomerResponse) func(w http.ResponseWriter, req *http.Request) {
		return func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			path := req.URL.Path
			arr := strings.Split(path, "/")
			version := arr[2]
			if !strings.Contains(req.RequestURI, version) {
				version = "v1"
			}
			slog.Debug("版本", slog.String("version", version))
			resp := apphandler(req, w)
			jsonrender.JSON(w, http.StatusOK, resp)
			Duration := time.Since(start)
			slog.Debug("received", slog.String("接收到请求的时间", start.Format("2006-01-02 15:04:05")), slog.String("查询用时", Duration.String()), slog.String("请求主机", req.Host), slog.String("请求方法", req.Method), slog.String("请求路径", req.URL.Path))
		}
	}

	router := mux.NewRouter()
	// http://127.0.0.1:9090/api/v1/getPersion?name=zen
	router.HandleFunc(url_prefix+"/v1/GetPersion", wrapper(controller.GetPersionInfo))

	router.HandleFunc(url_prefix+"/v1/GetWeathe", wrapper(controller.GetWeather))
	router.HandleFunc(url_prefix+"/v1/GetCity", wrapper(controller.GetCity))
	router.HandleFunc(url_prefix+"/v1/DeleteAllLive", wrapper(controller.DeleteAllLive))
	return router
}
