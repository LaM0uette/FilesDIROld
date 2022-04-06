package task

import (
	"FilesDIR/globals"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var (
	wg   sync.WaitGroup
	jobs chan string = make(chan string)
	Id               = 0
)

func loopFilesWorker() error {
	for path := range jobs {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			wg.Done()
			return err
		}

		for _, file := range files {
			if !file.IsDir() {
				fmt.Println(file.Name())
			}
		}
		wg.Done()
	}
	return nil
}

func LoopDirsFiles(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	go func() {
		wg.Add(1)
		jobs <- path
	}()
	for _, file := range files {
		if file.IsDir() {
			LoopDirsFiles(filepath.Join(path, file.Name()))
		}
	}
	return nil
}

func Run() {

	for w := 1; w <= 10; w++ {
		go loopFilesWorker()
	}

	LoopDirsFiles(globals.SrcPath)
	wg.Wait()
}
