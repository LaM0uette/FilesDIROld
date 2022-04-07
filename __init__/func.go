package __init__

import (
	"FilesDIR/globals"
	"os"
	"os/user"
	"path/filepath"
)

func init() {
	temp, err := user.Current()
	if err != nil {
		return
	}

	mainDir := filepath.Join(temp.HomeDir, globals.Name)
	logDir := filepath.Join(mainDir, "logs")
	dumpDir := filepath.Join(mainDir, "dumps")

	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}

	err = os.MkdirAll(dumpDir, os.ModePerm)
	if err != nil {
		return
	}
}
