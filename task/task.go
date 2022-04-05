package task

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func LoopDir(path string) error {

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			go func() {
				err := loopFiles(path)
				if err != nil {
					log.Println(err.Error())
				}
			}()
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func loopFiles(path string) error {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name())
		}
	}
	return nil
}

// Function qui pour chaque nouveau dossier lance une goroutine qui boucle sur les fichiers.
