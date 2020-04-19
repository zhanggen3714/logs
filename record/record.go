package record

import (
	"fmt"
	"hello/logs/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

//换新文件触发器
var (
	timeTriger bool
)

//FileRecord 日志记录到文件
type FileRecord struct {
	level  utils.LogLevel
	path   string
	config map[string]interface{}
	file   *os.File
	loginit *os.File
}

//NewLoger Loger的初始化函数(初始化时用户需要指定日志级别 )
func NewLoger(level, path string, config map[string]interface{}) *FileRecord {
	return &FileRecord{
		level:  utils.PaserLevel(level),
		path:   path,
		config: config,
	}

}

//判断文件夹是否存在
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {

		return false

	}
	return s.IsDir()

}

//判断文件是否存在
func isFileExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}


func(F FileRecord)updateLastRecordTime(){
	initFile := filepath.Join(F.path, "loginit")
	currentTimeString := time.Now().Format("20060102150405")
	Filestatus,_ := os.Stat(initFile)
	startOfEnd := Filestatus.Size() - 14
	fileObj, _ := os.OpenFile(initFile, os.O_RDWR, 744)
	fileObj.WriteAt([]byte(currentTimeString), startOfEnd)
	fileObj.WriteString(currentTimeString)
	fileObj.Close()
}

//加载文件末尾
func (F FileRecord) loaldloginitFileEnd() time.Time {
	initFile := filepath.Join(F.path, "loginit")
	currentTimeString := time.Now().Format("20060102150405")
	var bag = make([]byte, 14, 14)
	fileObj, _ := os.OpenFile(initFile, os.O_RDWR, 744)
	//读取文件的末尾
	Filestatus, err := os.Stat(initFile)
	startOfEnd := Filestatus.Size() - 14
	_, err = fileObj.ReadAt(bag, startOfEnd)
	if err == nil {
		fmt.Printf("%s\n", bag)
	}
	lastRecordTime, _ := time.Parse("20060102150405", string(bag))
	//最后1行解析时间失败重写
	if (lastRecordTime.Unix() >= 0) == false {
		fileObj.WriteAt([]byte(currentTimeString), startOfEnd)
		fileObj.WriteString(currentTimeString)
	}
	fileObj.Close()
	return lastRecordTime

}

func (F FileRecord) loaldloginitFile() {
	now := time.Now()
	LogPathExist := isDir(F.path)
	//检查日志路径是否存在？
	if LogPathExist == false {
		os.Mkdir(F.path, os.ModePerm)
	}
	//检查loginit文件是否存在
	initFile := filepath.Join(F.path, "loginit")
	currentTimeString := now.Format("20060102150405")
	if isFileExist(initFile) == false {
		os.Create(initFile)
		ioutil.WriteFile(initFile, []byte(currentTimeString), 744)
	}

	//获取loginit文件内容的第1行（首次加载log系统的时间）
	fileObj, _ := os.OpenFile(initFile, os.O_RDWR, 744)
	fileObj.Seek(0, 0)
	var bag = make([]byte, 14, 14)
	bag = bag[:]
	fileObj.Read(bag)
	CurrentInitFileContent := string(bag)
	InitalTime, _ := time.Parse("20060102150405", string(CurrentInitFileContent))
	//解析时间失败重写
	if (InitalTime.Unix() >= 0) == false {
		fileObj.Seek(0, 0)
		fileObj.WriteString(currentTimeString)
	}
	fileObj.Close()

}

//Duration 策略入口
func (F *FileRecord) Duration(during time.Duration) {
	now := time.Now()
	F.loaldloginitFile()
	lastRecordTime := F.loaldloginitFileEnd()
	NextFile := filepath.Join(F.path, lastRecordTime.Add(during).Format("20060102150405"))
	if now.After(lastRecordTime.Add(during)) {
		file, _ := os.Create(NextFile)
		F.file = file
		F.updateLastRecordTime()
		return
	}
	f, _ := os.OpenFile(NextFile,os.O_CREATE|os.O_APPEND, 744)
	F.file =f
}

//判断日志文件 迁移的标准 （时间|文件大小）
func (F *FileRecord) implemntStandard() {
	for k, v := range F.config {
		// //1.获取到结构体类型变量的反射类型
		refectCutStandard := reflect.ValueOf(F)
		// //2.获取确切的方法名
		method := refectCutStandard.MethodByName(k)
		// //3.带参数调用方式:构造一个类型为reflect.Value的切片
		args := []reflect.Value{
			reflect.ValueOf(v),
		}
		method.Call(args) //返回Value类型cls
	}

}

//Debug 记录debug级别日志信息
func (F *FileRecord) Debug(message string, a ...interface{}) {
	if F.level <= utils.DebugLevel {
		logs := utils.LogFormat(message, a...)
		F.implemntStandard()
		F.file.WriteString(logs)
		F.file.Close()
	}

}

//Info 方法记录Infor级别信息
func (F *FileRecord) Info(message string, a ...interface{}) {
	if F.level <= utils.InfoLevel {
		logs := utils.LogFormat(message, a...)
		F.implemntStandard()
		F.file.WriteString(logs)
		F.file.Close()
	}
}

//Erro 方法记录Erro级别信息
func (F *FileRecord) Erro(message string, a ...interface{}) {
	if F.level <= utils.ErroLevel {
		logs := utils.LogFormat(message, a...)
		F.implemntStandard()
		F.file.WriteString(logs)
		F.file.Close()
	}
}
