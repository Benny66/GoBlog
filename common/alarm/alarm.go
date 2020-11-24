package alarm

import (
	"encoding/json"
	"log"
	_func "goBlog/common/func"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type errorString struct {
	s string
}

type errorInfo struct {
	Time     string `json:"time"`
	Alarm    string `json:"alarm"`
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Funcname string `json:"funcname"`
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	go alarm("info", text, 2)
	return &errorString{text}
}

// 发邮件
func Email(text string) error {
	go alarm("email", text, 2)
	return &errorString{text}
}

// 发短信
func Sms(text string) error {
	go alarm("sms", text, 2)
	return &errorString{text}
}

// 发微信
func WeChat(text string) error {
	go alarm("wx", text, 2)
	return &errorString{text}
}

// Panic 异常
func Panic(text string) error {
	go alarm("error", text, 5)
	return &errorString{text}
}

func DDL(text string) error {
	go alarm("ddl", text, 6)
	return &errorString{text}
}

// 告警方法
func alarm(level string, str string, skip int) {
	// 当前时间
	currentTime := _func.GetTimeStr()

	// 定义 文件名、行号、方法名
	fileName, line, functionName := "?", 0, "?"

	pc, fileName, line, ok := runtime.Caller(skip)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}

	var msg = errorInfo{
		Time:     currentTime,
		Alarm:    level,
		Message:  str,
		Filename: fileName,
		Line:     line,
		Funcname: functionName,
	}

	jsons, errs := json.Marshal(msg)

	if errs != nil {
		log.Fatal(errs)
	}

	errorJsonInfo := string(jsons)

	var date = time.Now().Format("2006-01-02")
	logFilePath := _func.GetAbsPath("./runtime/logs/" + level + "-" + date + ".log")
	err := _func.IsFileExistsAndCreate(logFilePath)
	if err != nil {
		log.Fatal(err)
	}
	// 执行记日志
	f, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	// 设置日志输出到文件
	log.SetOutput(f)
	// 写入日志内容

	if level == "email" {
		// 执行发邮件

	} else if level == "sms" {
		// 执行发短信

	} else if level == "wx" {
		// 执行发微信

	} else if level == "info" {
		log.Println(errorJsonInfo)

	} else if level == "error" {
		// 执行PANIC方式
		log.Println(errorJsonInfo)
	} else if level == "ddl" {
		//发送
	}
}
