package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile(fmt.Sprintf("log/logs/SLog_%v.txt", time.Now().Format("20060102150405")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logDate := time.Now().Format("02/01/2006 15:04:05")

	Info = log.New(file, fmt.Sprintf("[-- INFO --][%v]: ", logDate), 0)
	Warning = log.New(file, fmt.Sprintf("[?? WARNING ??][%v]: ", logDate), log.Lshortfile)
	Error = log.New(file, fmt.Sprintf("[!! ERROR !!][%v]: ", logDate), log.Lshortfile)
}
