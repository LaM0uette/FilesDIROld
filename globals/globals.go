package globals

import (
	"log"
	"os/user"
	"path/filepath"
)

func tempPathGen() string {
	temp, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(temp.HomeDir, Name)
}

const (
	Name    = "FilesDIR"
	Author  = "LaM0uette"
	Version = "1.0.11"
)

var (
	TempPathGen   = tempPathGen()
	FolderLogs    = filepath.Join(tempPathGen(), "logs")
	FolderDumps   = filepath.Join(tempPathGen(), "dumps")
	FolderExports = filepath.Join(tempPathGen(), "exports")
)
