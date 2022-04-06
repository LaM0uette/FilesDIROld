package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Blank   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile(fmt.Sprintf("log/logs/SLog_%v.txt", time.Now().Format("20060102150405")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logDate := time.Now().Format("15:04:05")

	Blank = log.New(file, "", 0)
	Info = log.New(file, fmt.Sprintf("[%v][INFO]: ", logDate), 0)
	Warning = log.New(file, fmt.Sprintf("[%v][WARNING]: ", logDate), log.Lshortfile)
	Error = log.New(file, fmt.Sprintf("[%v][ERROR]: ", logDate), log.Lshortfile)
}
