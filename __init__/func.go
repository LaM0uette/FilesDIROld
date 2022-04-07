package __init__

import (
	"FilesDIR/globals"
	log2 "log"
	"os"
	"os/user"
	"path/filepath"
)

func mkdirFolder(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log2.Fatal(err)
	}
}

func init() {
	temp, err := user.Current()
	if err != nil {
		return
	}

	mainDir := filepath.Join(temp.HomeDir, globals.Name)
	logDir := filepath.Join(mainDir, "logs")
	dumpDir := filepath.Join(mainDir, "dumps")
	exportDir := filepath.Join(mainDir, "exports")

	mkdirFolder(logDir)
	mkdirFolder(dumpDir)
	mkdirFolder(exportDir)
}
