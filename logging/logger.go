package logging

import (
	"log"
	"os"
)

var Logger *log.Logger

const (
	DEBUG = "[DEBUG]"
	INFO  = "[INFO]"
	WARN  = "[WARN]"
	ERROR = "[ERROR]"
	FATAL = "[FATAL]"
)

func InitLog() {
	openFile, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Logger.Printf("[InitLog] open log file failed | err:%v", err)
	}
	Logger = log.New(openFile, "NOTE:", log.Lshortfile|log.LstdFlags)
}

func Debug(v ...interface{}) {
	Logger.SetPrefix(DEBUG)
	Logger.Println(v)
}

func Debugf(fmt string, v ...interface{}) {
	Logger.SetPrefix(DEBUG)
	Logger.Printf(fmt, v)
}

func Info(v ...interface{}) {
	Logger.SetPrefix(INFO)
	Logger.Println(v)
}

func Infof(fmt string, v ...interface{}) {
	Logger.SetPrefix(INFO)
	Logger.Printf(fmt, v)
}

func Warn(v ...interface{}) {
	Logger.SetPrefix(WARN)
	Logger.Println(v)
}

func Warnf(fmt string, v ...interface{}) {
	Logger.SetPrefix(WARN)
	Logger.Printf(fmt, v)
}

func Error(v ...interface{}) {
	Logger.SetPrefix(ERROR)
	Logger.Println(v)
}

func Errorf(fmt string, v ...interface{}) {
	Logger.SetPrefix(ERROR)
	Logger.Printf(fmt, v)
}

func Fatal(v ...interface{}) {
	Logger.SetPrefix(FATAL)
	Logger.Println(v)
}

func Fatalf(fmt string, v ...interface{}) {
	Logger.SetPrefix(FATAL)
	Logger.Printf(fmt, v)
}
