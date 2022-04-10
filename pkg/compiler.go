package pkg

import (
	"FilesDIR/globals"
	"os"
)

func ClsTempFiles() {

	_ = os.RemoveAll(globals.FolderLogs)
	_ = os.RemoveAll(globals.FolderDumps)
	_ = os.RemoveAll(globals.FolderExports)

	os.Exit(0)
}

func FicheAppuiFt() {
	return
}
