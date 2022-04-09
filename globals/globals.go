package globals

import (
	log2 "log"
	"os/user"
	"path/filepath"
)

func tempPathGen() string {
	temp, err := user.Current()
	if err != nil {
		log2.Fatal(err)
	}
	return filepath.Join(temp.HomeDir, Name)
}

const (
	Name    = "FilesDIR"
	Author  = "LaM0uette"
	Version = "1.0.0"
)

var (
	TempPathGen = tempPathGen()
)
