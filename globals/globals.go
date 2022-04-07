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
	Name       = "FilesDIR"
	Author     = "LaM0uette"
	Version    = 0.5
	SrcPathGen = "T:\\- 4 Suivi Appuis\\18-Partage\\de VILLELE DORIAN"
	DstPathGen = "T:\\- 4 Suivi Appuis\\26_MACROS\\GO\\FilesDIR\\Json\\json.json"
)

var (
	TempPathGen = tempPathGen()
)
