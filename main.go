package main

import (
	"hello/logs/record"
	"hello/logs/utils"
	"time"
)

//日志程序开启
func logsRecord(logger utils.Log) {
	for {
		time.Sleep(2 * time.Second)
		logger.Debug("这是一条debug 级别的日志信息")
		logger.Info("这是一条infor级别的日志信息")
		logger.Erro("这是一条Erro级别的日志信息")

	}

}

func main() {
	// logger := console.NewLoger("debug")
	config := make(map[string]interface{},1)
	config["Duration"] = time.Hour
	logger := record.NewLoger("debug", "D:\\goproject\\src\\hello\\logs\\logdb",config)
	// logger.Debug("这是一条debug 级别的日志信息")
	logsRecord(logger)

}
