package task

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var Id = 0

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
	fmt.Println("Finished", countDir, Id)
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
			go fmt.Println(file.Name())
			Id++
		}
	}

	wg.Done()
	return nil
}

func LoopDirsFiles(path string, wg *sync.WaitGroup) error {
	wg.Add(1)

	//fmt.Println(path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		wg.Done()
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name())
			//time.Sleep(1 * time.Second)
			Id++
		} else {
			go func() {
				err := LoopDirsFiles(filepath.Join(path, file.Name()), wg)
				if err != nil {
					log.Println(err)
				}
			}()
		}
	}

	wg.Done()
	return nil
}
