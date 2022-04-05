package task

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func LoopDir(path string) error {
	var wg sync.WaitGroup

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			wg.Add(1)

			go func() {
				err := loopFiles(path, &wg)
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

	wg.Wait()
	fmt.Println("Finished")
	return nil
}

func loopFiles(path string, wg *sync.WaitGroup) error {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name())
		}
	}

	wg.Done()
	return nil
}

// Function qui pour chaque nouveau dossier lance une goroutine qui boucle sur les fichiers.
