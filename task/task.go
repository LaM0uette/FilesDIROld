package task

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

var Id = 0
var Thr = 0

// LoopDir TODO: Code à supprimer
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

// loopFiles TODO: Code à supprimer
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

func SetProgramLimits() {
	const maxThreadCount int = 50 * 1000
	debug.SetMaxThreads(maxThreadCount)
}

func LoopDirsFiles(path string, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer func() { Thr-- }()
	defer wg.Done()

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && !strings.Contains(file.Name(), "~") {
			fmt.Println(file.Name(), Id, Thr)
			Id++
		} else if file.IsDir() {

			Thr++

			for Thr > 8000000 {

			}

			go func() {
				err = LoopDirsFiles(filepath.Join(path, file.Name()), wg)
				if err != nil {
					log.Print(err)
				}
			}()
			time.Sleep(20 * time.Millisecond)
		}
	}
	return nil
}
