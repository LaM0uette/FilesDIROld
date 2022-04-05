package task

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoopDir(path string) error {

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			fmt.Println("Dir: ", path)
		} else {
			fmt.Println("File : ", path)
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Function qui pour chaque nouveau dossier lance une goroutine qui boucle sur les fichiers.
