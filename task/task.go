package task

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var id = 0

func LoopDir(path string) error {
	var wg sync.WaitGroup

	countDir := 0

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			wg.Add(1)
			countDir++

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
	fmt.Println("Finished", countDir, id)
	return nil
}

func loopFiles(path string, wg *sync.WaitGroup) error {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		wg.Done()
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name())
			id++
		}
	}

	wg.Done()
	return nil
}
