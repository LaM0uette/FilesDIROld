package task

import (
	"FilesDIR/dump"
	"FilesDIR/log"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
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
	DstPath     string
	PoolSize    int
	NbFiles     int
	NbGoroutine int
}

func TempDir() string {
	temp, err := user.Current()
	if err != nil {
		log.Crash.Println("Error with your User folder.")
	}

	mainDir := filepath.Join(temp.HomeDir, "FilesDIR")
	logDir := filepath.Join(mainDir, "logs")
	dumpDir := filepath.Join(mainDir, "dumps")

	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log.Crash.Printf(fmt.Sprintf("Impossible to create log folder in: %s", logDir))
	}

	err = os.MkdirAll(dumpDir, os.ModePerm)
	if err != nil {
		log.Crash.Printf(fmt.Sprintf("Impossible to create dump folder in: %s", dumpDir))
	}

	return mainDir
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

				dump.Semicolon.Printf(fmt.Sprintf("%v;%s;%s;%s;%s",
					s.NbFiles, file.Name(), file.ModTime().Format("02-01-2006 15:04:05"), path+"/"+file.Name(), path))

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
			//time.Sleep(20 * time.Millisecond)
			LoopDirsFiles(filepath.Join(path, file.Name()))
		}
	}
}

func RunSearch(s *Sch) {

	DrawSetupSearch()

	dump.Semicolon.Println("id;Fichier;Date;Lien_Fichier;Lien")

	if s.PoolSize < 2 {
		log.Info.Println("Set the PoolSize to 2")
		s.PoolSize = 2
	}
	maxThr := s.PoolSize * 500

	log.Info.Printf(fmt.Sprintf("Set max thread count to %v\n\n", maxThr))
	debug.SetMaxThreads(maxThr)

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

	time.Sleep(1 * time.Second)

	DrawEndSearch()
}
