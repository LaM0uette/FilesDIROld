package __init__

import (
	"FilesDIR/globals"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func mkdirFolder(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func createFile(file string) {
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		var file, err = os.Create(file)
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)
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
	blacklistDir := filepath.Join(mainDir, "blacklist")

	mkdirFolder(logDir)
	mkdirFolder(dumpDir)
	mkdirFolder(exportDir)
	mkdirFolder(blacklistDir)

	createFile(filepath.Join(blacklistDir, "__ALL__.txt"))
}
