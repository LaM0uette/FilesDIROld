package task

import (
	"FilesDIR/dump"
	"FilesDIR/globals"
	"FilesDIR/log"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

type Flags struct {
	FlgMode      string
	FlgWord      string
	FlgExt       string
	FlgMaj       bool
	FlgXl        bool
	FlgDevil     bool
	FlgSuper     bool
	FlgBlackList bool
}

type Sch struct {
	SrcPath      string
	DstPath      string
	PoolSize     int
	NbFiles      int
	NbFilesTotal int
	NbGoroutine  int

	Mode      string
	Word      string
	BlackList []string
	Ext       string
	Maj       bool

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
	Wb          = &excelize.File{}
)

//...
// ACTIONS:
func strToLower(s string) string {
	return strings.ToLower(s)
}

func (s *Sch) getBlackList(file string) {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Crash.Println(err)
		os.Exit(1)
	}
	s.BlackList = append(s.BlackList, strings.Split(string(fileBytes), "\n")...)
}

func (s *Sch) checkFileSearched(file string) bool {
	name := file[:strings.LastIndex(file, path.Ext(file))]
	ext := strToLower(filepath.Ext(file))

	if !s.Maj {
		name = strToLower(name)
	}

	// condition of search Mode ( = | % | ^ | $ )
	switch s.Mode {
	case "%":
		if !strings.Contains(name, s.Word) {
			return false
		}
	case "=":
		if name != s.Word {
			return false
		}
	case "^":
		if !strings.HasPrefix(name, s.Word) {
			return false
		}
	case "$":
		if !strings.HasSuffix(name, s.Word) {
			return false
		}
	default:
		if !strings.Contains(name, s.Word) {
			return false
		}
	}

	// condition of extension file
	if s.Ext != ".*" && ext != s.Ext {
		return false
	}

	return true
}

func (s *Sch) isInBlackList(folderName string) bool {
	for _, black := range s.BlackList {
		if strings.Contains(folderName, black) {
			return true
		}
	}
	return false
}

//...
// WORKER:
func (s *Sch) loopFilesWorker(super bool) error {
	for pth := range jobs {
		files, err := ioutil.ReadDir(pth)
		if err != nil {
			log.Crash.Printf(fmt.Sprintf("Crash with this path: %s\n\n", pth))
			wg.Done()
			return err
		}

		for _, file := range files {

			if !file.IsDir() {
				if s.checkFileSearched(file.Name()) {
					s.NbFiles++
					s.NbFilesTotal++

					if !super {
						fmt.Print("\033[u\033[K", fmt.Sprintf("N°%v | Files: %s\n", s.NbFiles, file.Name()))

						log.BlankDate.Printf(fmt.Sprintf("N°%v | Files: %s", s.NbFiles, file.Name()))
						dump.Semicolon.Printf(fmt.Sprintf("%v;%s;%s;%s;%s",
							s.NbFiles, file.Name(), file.ModTime().Format("02-01-2006 15:04:05"), filepath.Join(pth, file.Name()), pth))
					}

					dataExp := exportData{
						Id:       s.NbFiles,
						File:     file.Name(),
						Date:     file.ModTime().Format("02-01-2006 15:04:05"),
						PathFile: filepath.Join(pth, file.Name()),
						Path:     pth,
					}
					ExcelData = append(ExcelData, dataExp)

					if runtime.NumGoroutine() > s.NbGoroutine {
						s.NbGoroutine = runtime.NumGoroutine()
					}
				} else {
					s.NbFilesTotal++
					fmt.Print("\033[u\033[K", fmt.Sprintf("Nombres de fichiers traités: %v", s.NbFilesTotal))
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

		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", job+2), ExcelData[job].Id)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", job+2), ExcelData[job].File)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", job+2), ExcelData[job].Date)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", job+2), ExcelData[job].PathFile)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", job+2), ExcelData[job].Path)

		wgWritter.Done()
	}
}

//...
// MAIN FUNC:
func (s *Sch) LoopDirsFiles(path string, f *Flags) {
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
		if file.IsDir() && !s.isInBlackList(file.Name()) {
			if f.FlgDevil {
				time.Sleep(20 * time.Millisecond)
				go s.LoopDirsFiles(filepath.Join(path, file.Name()), f)
			} else {
				s.LoopDirsFiles(filepath.Join(path, file.Name()), f)
			}
		}
	}
}

func RunSearch(s *Sch, f *Flags) {

	if !f.FlgSuper {
		log.BlankDate.Print(DrawInitSearch())
		fmt.Print(DrawInitSearch())
		time.Sleep(400 * time.Millisecond)

		if f.FlgBlackList {
			s.getBlackList(filepath.Join(globals.TempPathGen, "blacklist", "__ALL__.txt"))
		}
	}

	s.Mode = f.FlgMode
	s.Word = f.FlgWord
	if !f.FlgMaj {
		s.Word = strToLower(f.FlgWord)
	}
	s.Ext = fmt.Sprintf(".%s", f.FlgExt)
	s.Maj = f.FlgMaj

	if !f.FlgXl {
		dump.Semicolon.Println("id;Fichier;Date;Lien_Fichier;Lien")
		Wb = excelize.NewFile()
		_ = Wb.SetCellValue("Sheet1", "A1", "id")
		_ = Wb.SetCellValue("Sheet1", "B1", "Fichier")
		_ = Wb.SetCellValue("Sheet1", "C1", "Date")
		_ = Wb.SetCellValue("Sheet1", "D1", "LienFichier")
		_ = Wb.SetCellValue("Sheet1", "E1", "Lien")
	}

	if s.PoolSize < 2 {
		log.Info.Println("Set the PoolSize to 2")
		fmt.Println("Set the PoolSize to 2")
		s.PoolSize = 2
	}
	maxThr := s.PoolSize * 500

	if !f.FlgSuper {
		log.Info.Printf(fmt.Sprintf("Set max thread count to %v\n\n", maxThr))
		fmt.Printf("Set max thread count to %v\n\n", maxThr)
	}

	debug.SetMaxThreads(maxThr)

	searchStart := time.Now()

	fmt.Print("\033[s")
	for w := 1; w <= s.PoolSize; w++ {
		go func() {
			err := s.loopFilesWorker(f.FlgSuper)
			if err != nil {
				log.Error.Println(err)
			}
		}()
	}

	if !f.FlgSuper {
		log.Blank.Print(DrawRunSearch())
		fmt.Print(DrawRunSearch())
		time.Sleep(400 * time.Millisecond)
	}

	s.LoopDirsFiles(s.SrcPath, f)

	wg.Wait()

	s.TimerSearch = time.Since(searchStart)

	if !f.FlgSuper {
		time.Sleep(1 * time.Second)
		fmt.Print("\033[u\033[K")
		log.Blank.Print(DrawEndSearch())
		fmt.Print(DrawEndSearch())
		time.Sleep(200 * time.Millisecond)
	}

	// Export Excel
	if !f.FlgXl {
		if !f.FlgSuper {
			log.Blank.Print(DrawWriteExcel())
			fmt.Print(DrawWriteExcel())
		}

		fmt.Print("\033[s")

		iMax := len(ExcelData)
		for w := 1; w <= 500; w++ {
			go writeExcelLineWorker(Wb, iMax)
		}

		for i := 0; i < iMax-1; i++ {
			i := i
			go func() {
				wgWritter.Add(1)
				jobsWritter <- i
			}()
		}

		wgWritter.Wait()

		saveWord := f.FlgWord
		if len(f.FlgWord) < 1 {
			saveWord = "Export"
		}

		fmt.Print("\033[u\033[K")
		if err := Wb.SaveAs(filepath.Join(s.DstPath, saveWord+fmt.Sprintf("_%v.xlsx", time.Now().Format("20060102150405")))); err != nil {
			fmt.Println(err)
		}

		if !f.FlgSuper {
			log.Blank.Print(DrawSaveExcel())
			fmt.Println()
			fmt.Print(DrawSaveExcel())
			time.Sleep(600 * time.Millisecond)
		}
	}
}
