package task

import (
	"FilesDIR/log"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	jobs = make(chan string)
)

type Sch struct {
	SrcPath     string
	PoolSize    int
	NbFiles     int
	NbGoroutine int
}

func (s *Sch) loopFilesWorker() error {
	for path := range jobs {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Crash.Printf(fmt.Sprintf("Crash with this path: %s\n\n", path))
			wg.Done()
			return err
		}

		for _, file := range files {
			if !file.IsDir() {
				s.NbFiles++
				log.BlankDate.Printf(fmt.Sprintf("N°%v | Files: %s\n", s.NbFiles, file.Name()))
				fmt.Printf("N°%v | Files: %s\n", s.NbFiles, file.Name())

				if runtime.NumGoroutine() > s.NbGoroutine {
					s.NbGoroutine = runtime.NumGoroutine()
				}
			}
		}
		wg.Done()
	}
	return nil
}

func LoopDirsFiles(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Error.Printf(fmt.Sprintf("Error with this path: %s\n\n", path))
		fmt.Printf("Error with this path: %s\n\n", path)
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

func RunSearch(s *Sch) {

	if s.PoolSize < 2 {
		log.Info.Println("Set the PoolSize to 2\n")
		s.PoolSize = 2
	}

	maxThr := s.PoolSize + 500
	log.Info.Printf(fmt.Sprintf("Set max thread count to %v\n\n", maxThr))
	debug.SetMaxThreads(maxThr)

	DrawSetupSearch()

	for w := 1; w <= s.PoolSize; w++ {
		go func() {
			err := s.loopFilesWorker()
			if err != nil {
				log.Error.Println(err)
			}
		}()
	}

	DrawRunSearch()

	LoopDirsFiles(s.SrcPath)
	wg.Wait()

	time.Sleep(2 * time.Second)

	DrawEndSearch()
}
