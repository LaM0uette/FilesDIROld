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
	Cyan    = color.New(color.FgCyan).SprintFunc()
	Green   = color.New(color.FgHiGreen).SprintFunc()
	Red     = color.New(color.FgHiRed).SprintFunc()
	TRed    = color.New(color.FgRed).SprintFunc()
	Majenta = color.New(color.FgMagenta).SprintFunc()
	Black   = color.New(color.FgBlack).SprintFunc()
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
func Ui(v ...any) {
	ui.Print(v...)
	fmt.Print(Cyan(v...))
}

func Param(v ...any) {
	param.Print(v...)
	fmt.Print(Yellow(v...))
}

func Ok(v ...any) {
	ok.Print(v...)
	fmt.Print(Green(v...))
}

func Action(v ...any) {
	action.Print(v...)
	fmt.Print(Majenta(v...))
}

func End(v ...any) {
	ui.Print(v...)
	fmt.Print(Black(v...))
}

func Warning(v ...any) {
	warning.Print(v...)
	fmt.Print(TRed(v...))
}

func Error(v ...any) {
	errr.Print(v...)
	fmt.Print(Red(v...))
}

func Crash(v ...any) {
	crash.Print(v...)
	fmt.Print(Red(v...))
	os.Exit(1)
}

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
	action.Println(v...)
	fmt.Println(Majenta(v...))
}

func Endln(v ...any) {
	ui.Println(v...)
	fmt.Println(Black(v...))
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
	ok.Print(v...)
}

//...
// Print only
func POOk(v ...any) {
	fmt.Print(Green(v...))
}

func POAction(v ...any) {
	fmt.Print(Majenta(v...))
}

//...
// Dump
func Semicolon(v ...any) {
	vSemicolon.Println(v...)
}
