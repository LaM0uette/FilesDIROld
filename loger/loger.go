package loger

import (
	"FilesDIROLD/globals"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	// logs
	ui *log.Logger
)

func init() {
	logFile, err := os.OpenFile(filepath.Join(globals.FolderLogs, fmt.Sprintf("SLog_%v.txt", time.Now().Format("20060102150405"))), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ui = log.New(logFile, "", 0)
}

//...
// Log
func Start(v ...any) {
	ui.Print(v...)
}
