package utils

import (
	"fmt"
	"runtime"
	"time"
)

//Log 接口规范
type Log interface {
	Debug(message string, a ...interface{})
	Info(message string, a ...interface{})
	Erro(message string, a ...interface{})
}

//LogLevel 一个自定义类型(基于uint8)
type LogLevel uint8

//日志级别
const (
	DebugLevel LogLevel = iota
	InfoLevel
	ErroLevel
)

//日志切割标准编号
const (
	duration uint8 = iota
	size
)

//TraceError 记录执行调用关系
func TraceError(skip int) (functionName, filePath string, lineNumber int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller faild.error\n")
		return
	}
	functionName = runtime.FuncForPC(pc).Name()
	filePath = file
	lineNumber = line
	return
}

//PaserLevel 方法把用户输入的字符串level解析为自定义的LogLevel类型
func PaserLevel(level string) LogLevel {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "erro":
		return ErroLevel
	default:
		return DebugLevel
	}
}

//FileMagrationStandard 日志文件迁移的标准
type FileMagrationStandard struct {
	Standard map[string]interface{}
}


//LogFormat 日志格式处理
func LogFormat(message string, a ...interface{}) string {
	now := time.Now().Format("2006-01-02 15:04:05")
	functionName, file, line := TraceError(3)
	message = fmt.Sprintf(message, a...)
	erroMessage := fmt.Sprintf("%s %s %d", functionName, file, line)
	message = fmt.Sprintf("[ %s ] [ debug ] [%s] %s\n", now, erroMessage, message)
	return message
}
