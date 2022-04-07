package task

import (
	"FilesDIR/dump"
	"FilesDIR/log"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

type Sch struct {
	SrcPath     string
	DstPath     string
	PoolSize    int
	NbFiles     int
	NbGoroutine int
}

type exportData struct {
	Id       int    `json:"id"`
	File     string `json:"Fichier"`
	Date     string `json:"Date"`
	PathFile string `json:"Lien_Fichier"`
	Path     string `json:"Lien"`
}

var (
	wg        sync.WaitGroup
	jobs      = make(chan string)
	ExcelData []exportData
)

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

				dataExp := exportData{
					Id:       s.NbFiles,
					File:     file.Name(),
					Date:     file.ModTime().Format("02-01-2006 15:04:05"),
					PathFile: path + "/" + file.Name(),
					Path:     path,
				}
				ExcelData = append(ExcelData, dataExp)

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

	wb := excelize.NewFile()
	_ = wb.SetCellValue("Sheet1", "A1", "id")
	_ = wb.SetCellValue("Sheet1", "B1", "Fichier")
	_ = wb.SetCellValue("Sheet1", "C1", "Date")
	_ = wb.SetCellValue("Sheet1", "D1", "LienFichier")
	_ = wb.SetCellValue("Sheet1", "E1", "Lien")

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

	if err := wb.SaveAs(filepath.Join(s.DstPath, "word.xlsx")); err != nil {
		fmt.Println(err)
	}
}
