package loger

import (
	"FilesDIR/rgb"
	"FilesDIROLD/globals"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	// logs
	ui   *log.Logger
	errr *log.Logger
)

const (
	preErrr = "[ERROR]"
)

func init() {
	logFile, err := os.OpenFile(filepath.Join(globals.FolderLogs, fmt.Sprintf("SLog_%v.log", time.Now().Format("20060102150405"))), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ui = log.New(logFile, "", 0)
	errr = log.New(logFile, preErrr+" ", log.Ltime|log.Lmsgprefix)
}

//...
// Log
func Ui(v ...any) {
	ui.Print(v...)
}

func Error(msg string, err any) {
	errr.Print(msg, " ", err)
	fmt.Print(rgb.RedBg.Sprint(preErrr), rgb.RedB.Sprint(" ", msg), rgb.RedB.Sprint(" ", err), "\n")
}
