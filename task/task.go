package task

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var (
	wg   sync.WaitGroup
	jobs = make(chan string)
	Id   = 1
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
				Id++
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
			err := LoopDirsFiles(filepath.Join(path, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func RunSearch(path string, poolsize int) error {
	for w := 1; w <= poolsize; w++ {
		go func() {
			err := loopFilesWorker()
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	err := LoopDirsFiles(path)
	if err != nil {
		return err
	}
	wg.Wait()
	return nil
}
