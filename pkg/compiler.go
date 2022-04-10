package pkg

import (
	"FilesDIR/globals"
	"os"
)

func ClsTempFiles() {
	_ = os.RemoveAll(globals.FolderLogs)
	_ = os.RemoveAll(globals.FolderDumps)
	_ = os.RemoveAll(globals.FolderExports)
}

func FicheAppuiFt() {
	return
}
