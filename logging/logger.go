package logging

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLog() {
	openFile, err := os.OpenFile("./log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		Logger.Printf("[InitLog] open log file failed | err:%v", err)
	}
	Logger = log.New(openFile, "NOTE:", log.Lshortfile|log.LstdFlags)
}
