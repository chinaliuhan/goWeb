package main

import (
	"web/logs"
)

func main() {
	addr := "asdfaf"
	logs.Logger.Info("Start server at:%v", addr)
	err := "错误内容"
	logs.Logger.Critical("Server err:%v", err)
}
