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
	"strings"
	"sync"
	"time"
)

type Flags struct {
	FlgDevil bool
	FlgMode  string
	FlgWord  string
	FlgExt   string
	FlgMaj   bool
	FlgXl    bool
}

type Sch struct {
	SrcPath     string
	DstPath     string
	PoolSize    int
	NbFiles     int
	NbGoroutine int

	Mode string
	Word string
	Ext  string
	Maj  bool

	TimerSearch time.Duration
}

type exportData struct {
	Id       int    `json:"id"`
	File     string `json:"Fichier"`
	Date     string `json:"Date"`
	PathFile string `json:"Lien_Fichier"`
	Path     string `json:"Lien"`
}

var (
	wg          sync.WaitGroup
	wgWritter   sync.WaitGroup
	jobs        = make(chan string)
	jobsWritter = make(chan int)
	ExcelData   []exportData
)

//...
// ACTIONS:
func strToLower(s string, b bool) string {
	if !b {
		return strings.ToLower(s)
	} else {
		return s
	}
}

//...
// WORKER:
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

func writeExcelLineWorker(Wb *excelize.File, iMax int) {
	for job := range jobsWritter {
		fmt.Print("\033[u\033[K")
		fmt.Printf("%v/%v", job, iMax)

		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", job), ExcelData[job].Id)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", job), ExcelData[job].File)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", job), ExcelData[job].Date)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", job), ExcelData[job].PathFile)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", job), ExcelData[job].Path)

		wgWritter.Done()
	}
}

//...
// MAIN FUNC:
func LoopDirsFiles(path string, f *Flags) {
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
			if f.FlgDevil {
				time.Sleep(20 * time.Millisecond)
				go LoopDirsFiles(filepath.Join(path, file.Name()), f)
			} else {
				LoopDirsFiles(filepath.Join(path, file.Name()), f)
			}
		}
	}
}

func RunSearch(s *Sch, f *Flags) {

	DrawSetupSearch()

	s.Mode = f.FlgMode
	s.Word = strToLower(s.Word, s.Maj)
	s.Ext = f.FlgExt
	s.Maj = f.FlgMaj

	dump.Semicolon.Println("id;Fichier;Date;Lien_Fichier;Lien")
	Wb := excelize.NewFile()
	_ = Wb.SetCellValue("Sheet1", "A1", "id")
	_ = Wb.SetCellValue("Sheet1", "B1", "Fichier")
	_ = Wb.SetCellValue("Sheet1", "C1", "Date")
	_ = Wb.SetCellValue("Sheet1", "D1", "LienFichier")
	_ = Wb.SetCellValue("Sheet1", "E1", "Lien")

	if s.PoolSize < 2 {
		log.Info.Println("Set the PoolSize to 2")
		s.PoolSize = 2
	}
	maxThr := s.PoolSize * 500

	searchStart := time.Now()
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

	LoopDirsFiles(s.SrcPath, f)

	wg.Wait()
	s.TimerSearch = time.Since(searchStart)

	time.Sleep(1 * time.Second)

	DrawEndSearch()

	time.Sleep(200 * time.Millisecond)

	// Export Excel
	if !f.FlgXl {
		DrawWriteExcel()

		fmt.Print("\033[s")

		iMax := len(ExcelData)
		for w := 1; w <= 500; w++ {
			go writeExcelLineWorker(Wb, iMax)
		}

		for i := 1; i < iMax-1; i++ {
			i := i
			go func() {
				wgWritter.Add(1)
				jobsWritter <- i
			}()
		}

		wgWritter.Wait()

		fmt.Print("\033[u\033[K")
		if err := Wb.SaveAs(filepath.Join(s.DstPath, "word.xlsx")); err != nil {
			fmt.Println(err)
		}

		DrawSaveExcel()
		time.Sleep(600 * time.Millisecond)
	}
}
