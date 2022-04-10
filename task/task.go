package task

import (
	"FilesDIR/construct"
	"FilesDIR/globals"
	"FilesDIR/loger"
	"bufio"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Search struct {
	SrcPath      string
	DstPath      string
	NbFiles      int
	NbFilesTotal int
	NbGoroutine  int

	Mode      string
	Word      string
	BlackList []string
	Ext       string
	Maj       bool

	TimerSearch time.Duration
	ReqFinal    string
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
func CurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func strToLower(s string) string {
	return strings.ToLower(s)
}

func (s *Search) getBlackList(file string) {

	readFile, err := os.Open(file)
	if err != nil {
		loger.Crashln(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s.BlackList = append(s.BlackList, fileScanner.Text())
	}

	_ = readFile.Close()
}

func (s *Search) isInBlackList(folderName string) bool {
	for _, black := range s.BlackList {
		if strings.Contains(strToLower(folderName), strToLower(black)) {
			return true
		}
	}
	return false
}

func (s *Search) checkFileSearched(file string) bool {
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

//...
// WORKER:
func (s *Search) loopFilesWorker(super bool) error {
	for pth := range jobs {
		files, err := ioutil.ReadDir(pth)
		if err != nil {
			loger.Crashln(fmt.Sprintf("Crash with this path: %s", pth))
			wg.Done()
			return err
		}

		for _, file := range files {

			if !file.IsDir() {
				if s.checkFileSearched(file.Name()) {
					s.NbFiles++
					s.NbFilesTotal++

					if !super {
						fmt.Print(fmt.Sprintf("N°%v | Files: %s\n", s.NbFiles, file.Name()))

						dataExp := exportData{
							Id:       s.NbFiles,
							File:     file.Name(),
							Date:     file.ModTime().Format("02-01-2006 15:04:05"),
							PathFile: filepath.Join(pth, file.Name()),
							Path:     pth,
						}
						ExcelData = append(ExcelData, dataExp)

					} else {
						fmt.Print(fmt.Sprintf("\rNombres de fichiers traités: %v", s.NbFilesTotal))
					}

					loger.BlankDateln(fmt.Sprintf("N°%v | Files: %s", s.NbFiles, file.Name()))
					loger.Semicolonln(fmt.Sprintf("%v;%s;%s;%s;%s",
						s.NbFiles, file.Name(), file.ModTime().Format("02-01-2006 15:04:05"), filepath.Join(pth, file.Name()), pth))

					if runtime.NumGoroutine() > s.NbGoroutine {
						s.NbGoroutine = runtime.NumGoroutine()
					}
				} else {
					s.NbFilesTotal++
					fmt.Print(fmt.Sprintf("\nNombres de fichiers traités: %v", s.NbFilesTotal))
				}
			}
		}
		wg.Done()
	}
	return nil
}

func writeExcelLineWorker(Wb *excelize.File, iMax int) {

	for job := range jobsWritter {

		fmt.Printf("\rSauvegarde du fichier Excel...  %v/%v", job, iMax)

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
func (s *Search) LoopDirsFiles(path string, f *construct.Flags) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		loger.Errorln(fmt.Sprintf("Error with this path: %s", path))
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

func (s *Search) RunSearch(f *construct.Flags) {

	s.Mode = f.FlgMode
	s.Word = f.FlgWord
	if !f.FlgMaj {
		s.Word = strToLower(f.FlgWord)
	}
	s.Ext = fmt.Sprintf(".%s", f.FlgExt)
	s.Maj = f.FlgMaj
	s.ReqFinal = f.GetReqOfSearched()

	f.DrawInitSearch()

	// Check blacklist if enabled and add items in it
	if f.FlgBlackList {
		s.getBlackList(filepath.Join(globals.TempPathGen, "blacklist", "__ALL__.txt"))

		file := filepath.Join(globals.TempPathGen, "blacklist", fmt.Sprintf("%s.txt", strToLower(s.Word)))
		if _, err := os.Stat(file); err == nil {
			s.getBlackList(file)
		}
	}

	if f.ExportExcelActivate() {
		loger.Semicolonln("id;Fichier;Date;Lien_Fichier;Lien")

		Wb = excelize.NewFile()
		_ = Wb.SetCellValue("Sheet1", "A1", "id")
		_ = Wb.SetCellValue("Sheet1", "B1", "Fichier")
		_ = Wb.SetCellValue("Sheet1", "C1", "Date")
		_ = Wb.SetCellValue("Sheet1", "D1", "LienFichier")
		_ = Wb.SetCellValue("Sheet1", "E1", "Lien")
	}

	f.CheckMinimumPoolSize()

	f.SetMaxThread()

	f.DrawRunSearch()

	timeSearchStart := time.Now()

	// Creation of workers for search
	for w := 1; w <= f.FlgPoolSize; w++ {
		go func() {
			err := s.loopFilesWorker(f.FlgSuper)
			if err != nil {
				loger.Errorln(err)
			}
		}()
	}
	// Run search loop
	s.LoopDirsFiles(s.SrcPath, f)

	wg.Wait() // Wait for all search loops to complete

	s.TimerSearch = time.Since(timeSearchStart)

	f.DrawEndSearch()

	// Export xlsx
	if f.ExportExcelActivate() {

		f.DrawWriteExcel()

		// Creation of workers for write line in excel file
		iMax := len(ExcelData)
		for w := 1; w <= 300; w++ {
			go writeExcelLineWorker(Wb, iMax)
		}
		// Run writing loop
		for i := 0; i < iMax-1; i++ {
			i := i
			go func() {
				wgWritter.Add(1)
				jobsWritter <- i
			}()
		}

		wgWritter.Wait() // Wait for all write loops to complete

		// Generate a default word if is none
		saveWord := f.SetSaveWord()

		// Save Excel file
		if err := Wb.SaveAs(filepath.Join(s.DstPath, saveWord+fmt.Sprintf("_%v.xlsx", time.Now().Format("20060102150405")))); err != nil {
			fmt.Println(err)
		}

		f.DrawSaveExcel()
	}
}
