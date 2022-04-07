package dump

import (
	"FilesDIR/globals"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	Semicolon *log.Logger
)

func init() {
	file, err := os.OpenFile(filepath.Join(globals.TempPathGen, "dumps", fmt.Sprintf("Dump_%v.txt", time.Now().Format("20060102150405"))), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Semicolon = log.New(file, "", 0)
}
