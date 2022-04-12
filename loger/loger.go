package loger

import (
	"FilesDIR/globals"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	// logs
	ui         *log.Logger
	vBlankDate *log.Logger
	vInfo      *log.Logger
	vWarning   *log.Logger
	vError     *log.Logger
	crash      *log.Logger

	// dumps
	vSemicolon *log.Logger

	// colors
	Cyan = color.New(color.FgCyan).SprintFunc()
	Red  = color.New(color.FgRed).SprintFunc()
)

func init() {
	logFile, err := os.OpenFile(filepath.Join(globals.FolderLogs, fmt.Sprintf("SLog_%v.txt", time.Now().Format("20060102150405"))), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	dumpFile, err := os.OpenFile(filepath.Join(globals.FolderDumps, fmt.Sprintf("Dump_%v.txt", time.Now().Format("20060102150405"))), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ui = log.New(logFile, "", 0)
	vBlankDate = log.New(logFile, ": ", log.Ltime|log.Lmsgprefix)
	vInfo = log.New(logFile, "[INFO]: ", log.Ltime|log.Lmsgprefix)
	vWarning = log.New(logFile, "[WARNING]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)
	vError = log.New(logFile, "[ERROR]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)
	crash = log.New(logFile, "[CRASH]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)

	vSemicolon = log.New(dumpFile, "", 0)
}

//...
// Log
func Ui(v ...any) {
	ui.Println(v...)
	fmt.Println(Cyan(v...))
}

func Crash(v ...any) {
	crash.Println(v...)
	fmt.Println(Red(v...))
	os.Exit(1)
}

func Blank(v ...any) {
	ui.Print(v...)
	fmt.Print(v...)
}

func Blankln(v ...any) {
	ui.Println(v...)
	fmt.Println(v...)
}

func BlankDateln(v ...any) {
	vBlankDate.Println(v...)
	fmt.Println(v...)
}

func Infoln(v ...any) {
	vInfo.Println(v...)
	fmt.Println(v...)
}

func Warningln(v ...any) {
	vWarning.Println(v...)
	fmt.Println(v...)
}

func Errorln(v ...any) {
	vError.Println(v...)
	fmt.Println(v...)
}

//...
// Log only
func LOBlankDateln(v ...any) {
	vBlankDate.Println(v...)
}

//...
// Dump
func Semicolonln(v ...any) {
	vSemicolon.Println(v...)
}
