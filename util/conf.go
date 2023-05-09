package util

import (
	"github.com/zhangyiming748/goini"
	"golang.org/x/exp/slog"
)

const confPath = "./conf.ini"

var (
	RunMode string
	Conf    *goini.Config
)

/*
*
  - 初始化
    init函数的主要作用：
    初始化不能采用初始化表达式初始化的变量。
    程序运行前的注册。
    实现sync.Once功能。
    其他
*/
func init() {
	initConfig()
}

func initConfig() {
	Conf = goini.SetConfig(confPath)
	slog.Info(confPath)
}

/**
 * 获取环境变量
 */
func GetEnv() string {
	if RunMode == "" {
		initConfig()
	}
	return RunMode
}

/**
 * 根据键获取值
 */
func GetVal(section, name string) string {
	if section == "" {
		section = GetEnv()
	}
	val, _ := Conf.GetValue(section, name)
	return val
}
