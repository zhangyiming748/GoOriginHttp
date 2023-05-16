package main

import (
	"GoOriginHttp/api"
	"GoOriginHttp/controller"
	"GoOriginHttp/model"
	"GoOriginHttp/mysql"
	"GoOriginHttp/util"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	startAt := time.Now()
	ch := make(chan os.Signal)
	// 监听信号
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range ch {
			switch s { // 终端控制进程结束(终端连接断开)|用户发送INTR字符(Ctrl+C)触发|结束程序
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				//slog.Debug("退出服务", slog.Any("信号量", s))
				end := time.Now().Sub(startAt).Minutes()
				endAt := fmt.Sprintf("正常运行%.2f分后停止服务", end)
				slog.Info(endAt)
				os.Exit(0)
			default:
				slog.Debug("其他信号:", slog.Any("信号量", s))
			}
		}
	}()
	SetLog(util.GetVal("log", "level"))

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

	router.HandleFunc(url_prefix+"/v1/GetWeather", wrapper(controller.GetWeather))
	router.HandleFunc(url_prefix+"/v1/GetCity", wrapper(controller.GetCity))
	router.HandleFunc(url_prefix+"/v1/DeleteAllLive", wrapper(controller.DeleteAllLive))
	router.HandleFunc(url_prefix+"/v1/GetAllLive", wrapper(controller.GetAllLive))

	return router
}
