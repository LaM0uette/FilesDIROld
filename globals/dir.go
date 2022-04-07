package globals

import (
	log2 "log"
	"os"
	"os/user"
	"path/filepath"
)

func TempDir() string {
	temp, err := user.Current()
	if err != nil {
		log2.Fatal(err)
	}

	mainDir := filepath.Join(temp.HomeDir, "FilesDIR")
	logDir := filepath.Join(mainDir, "logs")
	dumpDir := filepath.Join(mainDir, "dumps")

	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log2.Fatal(err)
	}

	err = os.MkdirAll(dumpDir, os.ModePerm)
	if err != nil {
		log2.Fatal(err)
	}

	return mainDir
}
