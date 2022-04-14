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
	ui    *log.Logger
	ok    *log.Logger
	errr  *log.Logger
	crash *log.Logger

	semicolon *log.Logger
)

const (
	preOk    = "[ -OK- ]"
	preErrr  = "[ -ERROR- ]"
	preCrash = "[ -CRASH- ]"
)

func init() {
	logFile, err := os.OpenFile(filepath.Join(globals.FolderLogs, fmt.Sprintf("SLog_%v.log", time.Now().Format("20060102150405"))), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	dumpFile, err := os.OpenFile(filepath.Join(globals.FolderDumps, fmt.Sprintf("Dump_%v.csv", time.Now().Format("20060102150405"))), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ui = log.New(logFile, "", 0)
	ok = log.New(logFile, preOk+" ", log.Ltime|log.Lmsgprefix)
	errr = log.New(logFile, preErrr+" ", log.Ltime|log.Lmsgprefix)
	crash = log.New(logFile, preCrash+" ", log.Ltime|log.Lmsgprefix)

	semicolon = log.New(dumpFile, "", 0)
}

//...
// Log
func Ui(v ...any) {
	ui.Print(v...)
}

func Ok(msg string) {
	ok.Print(msg)
	fmt.Print(rgb.GreenBg.Sprint(preErrr), rgb.GreenB.Sprint(" ", msg))
}

func Error(msg string, err any) {
	errr.Print(msg, " ", err)
	fmt.Print(rgb.RedBg.Sprint(preErrr), rgb.RedB.Sprint(" ", msg), rgb.RedB.Sprint(" ", err), "\n")
}

func Crash(msg string, err any) {
	crash.Print(msg, " ", err)
	fmt.Print(rgb.RedBg.Sprint(preCrash), rgb.RedBg.Sprint(" ", msg), rgb.RedB.Sprint(" ", err), "\n")
	os.Exit(1)
}

//...
// Dump
func Semicolon(v ...any) {
	semicolon.Println(v...)
}
