package task

import (
	"FilesDIR/construct"
	"FilesDIR/globals"
	"FilesDIR/loger"
	"bufio"
	"fmt"
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

var (
	wg   sync.WaitGroup
	jobs = make(chan string)
	Mu   sync.Mutex
)

//...
// ACTIONS:
func CurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		loger.Errorln(err)
		os.Exit(1)
	}
	return pwd
}

func StrToLower(s string) string {
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
		if strings.Contains(StrToLower(folderName), StrToLower(black)) {
			return true
		}
	}
	return false
}

func (s *Search) checkFileSearched(file string) bool {
	name := file[:strings.LastIndex(file, path.Ext(file))]
	ext := StrToLower(filepath.Ext(file))

	if !s.Maj {
		name = StrToLower(name)
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

	// condition of open file
	if strings.Contains(name, "~") {
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
			loger.Crashln(fmt.Sprintf("Crashln with this path: %s", pth))
			wg.Done()
			return err
		}

		for _, file := range files {

			if !file.IsDir() {
				if s.checkFileSearched(file.Name()) {
					s.NbFiles++
					s.NbFilesTotal++

					if !super {
						Mu.Lock()
						loger.POOk(fmt.Sprintf("\rN°%v | Files: %s\n", s.NbFiles, file.Name()))

						dataExp := construct.ExportData{
							Id:       s.NbFiles,
							File:     file.Name(),
							Date:     file.ModTime().Format("02-01-2006 15:04:05"),
							PathFile: filepath.Join(pth, file.Name()),
							Path:     pth,
						}
						construct.ExcelData = append(construct.ExcelData, dataExp)
						Mu.Unlock()

					} else {
						loger.POAction(fmt.Sprintf("\rFait: %v", s.NbFilesTotal))
					}

					loger.LOOk(fmt.Sprintf("N°%v | Files: %s", s.NbFiles, file.Name()))
					loger.Semicolon(fmt.Sprintf("%v;%s;%s;%s;%s",
						s.NbFiles, file.Name(), file.ModTime().Format("02-01-2006 15:04:05"), filepath.Join(pth, file.Name()), pth))

					if runtime.NumGoroutine() > s.NbGoroutine {
						s.NbGoroutine = runtime.NumGoroutine()
					}
				} else {
					s.NbFilesTotal++
					loger.POAction(fmt.Sprintf("\rFait: %v", s.NbFilesTotal))
				}
			}
		}
		wg.Done()
	}
	return nil
}

//...
// MAIN FUNC:
func (s *Search) LoopDirsFiles(path string, f *construct.Flags) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// TODO A check
		loger.Errorln(fmt.Sprintf("<fg=158, 21, 3>Error with this path:</> <fg=230, 31, 5>%s</>", path))
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
		s.Word = StrToLower(f.FlgWord)
	}
	s.Ext = fmt.Sprintf(".%s", f.FlgExt)
	s.Maj = f.FlgMaj
	s.ReqFinal = f.GetReqOfSearched()

	f.DrawInitSearch()

	// Check blacklist if enabled and add items in it
	if f.FlgBlackList {
		s.getBlackList(filepath.Join(globals.TempPathGen, "blacklist", "__ALL__.txt"))

		file := filepath.Join(globals.TempPathGen, "blacklist", fmt.Sprintf("%s.txt", StrToLower(s.Word)))
		if _, err := os.Stat(file); err == nil {
			s.getBlackList(file)
		}
	}

	f.CheckMinimumPoolSize()

	f.SetMaxThread()

	// Generate column of dump
	loger.Semicolon("id;Fichier;Date;Lien_Fichier;Lien")

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
	f.GenerateExcelSave(s.DstPath)
}
