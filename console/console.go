package console

import (
	"fmt"
	"hello/logs/utils"
)

//Loger 日志对象
type Loger struct {
	level utils.LogLevel
}

//NewLoger Loger的初始化函数(初始化时用户需要指定日志级别 )
func NewLoger(level string) Loger {
	return Loger{
		level: utils.PaserLevel(level),
	}

}

//Debug 打印debug级别日志信息
func (L Loger) Debug(message string, a ...interface{}) {
	if L.level <= utils.DebugLevel {
		logs := utils.LogFormat(message, a...)
		fmt.Println(logs)
	}

}

//Info 方法打印Infor级别信息
func (L Loger) Info(message string, a ...interface{}) {
	if L.level <= utils.InfoLevel {
		logs := utils.LogFormat(message, a...)
		fmt.Println(logs)
	}
}

//Erro 方法打印Erro级别信息
func (L Loger) Erro(message string, a ...interface{}) {
	if L.level <= utils.ErroLevel {
		logs := utils.LogFormat(message, a...)
		fmt.Println(logs)
	}
}
