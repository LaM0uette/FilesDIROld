package loger

import (
	"FilesDIR/globals"
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	// logs
	ui      *log.Logger
	param   *log.Logger
	ok      *log.Logger
	action  *log.Logger
	warning *log.Logger
	errr    *log.Logger
	crash   *log.Logger

	// dumps
	vSemicolon *log.Logger

	// colors
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
	param = log.New(logFile, "[INFO]: ", log.Ltime|log.Lmsgprefix)
	ok = log.New(logFile, ": ", log.Ltime|log.Lmsgprefix)
	action = log.New(logFile, ": ", log.Ltime|log.Lmsgprefix)
	warning = log.New(logFile, "[WARNING]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)
	errr = log.New(logFile, "[ERROR]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)
	crash = log.New(logFile, "[CRASH]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)

	vSemicolon = log.New(dumpFile, "", 0)
}

//...
// Log
func Start(v ...any) {
	ui.Print(v...)
	color.HiCyan.Print(v...)
}

func Ui(v ...any) {
	ui.Print(v...)
	color.Print(v...)
}

func Param(v ...any) {
	param.Print(v...)
	color.Print(v...)
}

func Ok(v ...any) {
	ok.Print(v...)
	color.Print(v...)
}

func Action(v ...any) {
	action.Print(v...)
	color.Print(v...)
}

func End(v ...any) {
	ui.Print(v...)
	color.Print(v...)
}

func Warning(v ...any) {
	warning.Print(v...)
	color.Print(v...)
}

func Error(v ...any) {
	errr.Print(v...)
	color.Print(v...)
}

func Crash(v ...any) {
	crash.Print(v...)
	color.Print(v...)
	os.Exit(1)
}

func Uiln(v ...any) {
	ui.Println(v...)
	color.Println(v...)
}

func Paramln(v ...any) {
	param.Println(v...)
	color.Println(v...)
}

func Okln(v ...any) {
	ok.Println(v...)
	color.Println(v...)
}

func Actionln(v ...any) {
	action.Println(v...)
	color.Println(v...)
}

func Endln(v ...any) {
	ui.Println(v...)
	color.Println(v...)
}

func Warningln(v ...any) {
	warning.Println(v...)
	color.Println(v...)
}

func Errorln(v ...any) {
	errr.Println(v...)
	color.Println(v...)
}

func Crashln(v ...any) {
	crash.Println(v...)
	color.Println(v...)
	os.Exit(1)
}

//...
// Log only
func LOOk(v ...any) {
	ok.Print(v...)
}

//...
// Print only
func POOk(v ...any) {
	color.Print(v...)
}

func POAction(v ...any) {
	color.Print(v...)
}

//...
// Dump
func Semicolon(v ...any) {
	vSemicolon.Println(v...)
}
