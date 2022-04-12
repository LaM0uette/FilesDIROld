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
	Cyan    = color.New(color.FgHiCyan).SprintFunc()
	Green   = color.New(color.FgHiGreen).SprintFunc()
	Red     = color.New(color.FgHiRed).SprintFunc()
	TRed    = color.New(color.FgRed).SprintFunc()
	Majenta = color.New(color.FgHiMagenta).SprintFunc()
	Yellow  = color.New(color.FgHiYellow).SprintFunc()
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
func Uiln(v ...any) {
	ui.Println(v...)
	fmt.Println(Cyan(v...))
}

func Paramln(v ...any) {
	param.Println(v...)
	fmt.Println(Yellow(v...))
}

func Okln(v ...any) {
	ok.Println(v...)
	fmt.Println(Green(v...))
}

func Actionln(v ...any) {
	action.Print(v...)
	fmt.Print(Majenta(v...))
}

func Warningln(v ...any) {
	warning.Println(v...)
	fmt.Println(TRed(v...))
}

func Errorln(v ...any) {
	errr.Println(v...)
	fmt.Println(Red(v...))
}

func Crashln(v ...any) {
	crash.Println(v...)
	fmt.Println(Red(v...))
	os.Exit(1)
}

//...
// Log only
func LOOk(v ...any) {
	ok.Println(v...)
}

//...
// Print only
func POOk(v ...any) {
	fmt.Println(Green(v...))
}

//...
// Dump
func Semicolon(v ...any) {
	vSemicolon.Println(v...)
}
