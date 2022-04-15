package config

import (
	"log"
	"os/user"
	"path/filepath"
)

const (
	Name    = "FilesDIR"
	Author  = "LaM0uette"
	Version = "1.1.0"
)

var (
	DstPath = filepath.Join(GetTempDir(), Name)
)

func GetTempDir() string {
	temp, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(temp.HomeDir)
}
