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
	Version = "1.1.0"

	// RGB
	Th1 = "44,168,65"
	Th2 = "86,195,199"
	//Th3   = "86,195,199"
	Param = "196,168,27"
)

var (
	TempPathGen   = tempPathGen()
	FolderLogs    = filepath.Join(tempPathGen(), "logs")
	FolderDumps   = filepath.Join(tempPathGen(), "dumps")
	FolderExports = filepath.Join(tempPathGen(), "exports")
)
