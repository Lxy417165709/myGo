package main

import (
	"conf/env"
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.SetLogFuncCallDepth(3)
	logs.EnableFuncCallDepth(true)
}

func main() {
	const pathOfConf = "conf/conf_test.json"
	if err := env.Conf.Load(pathOfConf); err != nil {
		logs.Error(err)
		return
	}
	logs.Info("%+v",env.Conf)
}
