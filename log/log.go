package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	//Blank     *log.Logger
	BlankDate *log.Logger
	Info      *log.Logger
	//Warning   *log.Logger
	Error *log.Logger
	Crash *log.Logger
)

func init() {
	file, err := os.OpenFile(fmt.Sprintf("log/dumps/SLog_%v.txt", time.Now().Format("20060102150405")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	//Blank = log.New(file, "", 0)
	BlankDate = log.New(file, ": ", log.Ltime|log.Lmsgprefix)
	Info = log.New(file, "[INFO]: ", log.Ltime|log.Lmsgprefix)
	//Warning = log.New(file, "[WARNING]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)
	Error = log.New(file, "[ERROR]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)
	Crash = log.New(file, "[CRASH]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)
}
