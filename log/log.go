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
	//Info      *log.Logger
	//Warning   *log.Logger
	Error *log.Logger
)

func init() {
	file, err := os.OpenFile(fmt.Sprintf("log/logs/SLog_%v.txt", time.Now().Format("20060102150405")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	//Blank = log.New(file, "", 0)
	BlankDate = log.New(file, fmt.Sprintf("[%v]: ", log.Ltime), 0)
	//Info = log.New(file, fmt.Sprintf("[%v][INFO]: ", log.Ltime), 0)
	//Warning = log.New(file, fmt.Sprintf("[%v][WARNING]: ", log.Ltime), log.Lshortfile)
	Error = log.New(file, fmt.Sprintf("[%v][ERROR]: ", log.Ltime), log.Lshortfile)
}
