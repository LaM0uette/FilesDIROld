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
	Id   = 0
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
				Id++
				fmt.Println(file.Name())
			}
		}
		wg.Done()
	}
	return nil
}

func LoopDirsFiles(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Print(err)
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
}

func RunSearch(path string, poolsize int) {
	for w := 1; w <= poolsize; w++ {
		go func() {
			err := loopFilesWorker()
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	LoopDirsFiles(path)
	wg.Wait()
}
