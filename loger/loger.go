package loger

import (
	"FilesDIR/globals"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	vBlank     *log.Logger
	vBlankDate *log.Logger
	vInfo      *log.Logger
	vWarning   *log.Logger
	vError     *log.Logger
	vCrash     *log.Logger
)

func init() {
	file, err := os.OpenFile(filepath.Join(globals.TempPathGen, "logs", fmt.Sprintf("SLog_%v.txt", time.Now().Format("20060102150405"))), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	vBlank = log.New(file, "", 0)
	vBlankDate = log.New(file, ": ", log.Ltime|log.Lmsgprefix)
	vInfo = log.New(file, "[INFO]: ", log.Ltime|log.Lmsgprefix)
	vWarning = log.New(file, "[WARNING]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)
	vError = log.New(file, "[ERROR]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)
	vCrash = log.New(file, "[CRASH]: ", log.Ltime|log.Lmsgprefix|log.Lshortfile)
}

func Blankln(v ...any) {
	vBlank.Println(v...)
	fmt.Println(v...)
}

func Blank(v ...any) {
	vBlank.Print(v...)
	fmt.Print(v...)
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

func Crashln(v ...any) {
	vCrash.Println(v...)
	fmt.Println(v...)
	os.Exit(1)
}
