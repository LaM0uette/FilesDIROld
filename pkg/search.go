package pkg

import (
	"FilesDIR/loger"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Timer struct {
	AppStart time.Time
	AppEnd   time.Duration

	SearchStart time.Time
	SearchEnd   time.Duration
}

type Counter struct {
	NbrFiles    uint64
	NbrAllFiles uint64

	NbrFolder uint64
}

type Process struct {
	NbrThreads    int
	NbrGoroutines int
}

type Search struct {

	// Flags
	Cls      bool
	Compiler bool

	//..
	Mode      string
	Word      string
	Ext       string
	PoolSize  int
	Maj       bool
	Devil     bool
	Silent    bool
	WordsList bool
	BlackList bool
	WhiteList bool

	// Search
	SrcPath string
	DstPath string
	ReqUse  string

	// Data
	ListWords     []string
	ListBlackList []string
	ListWhiteList []string
	Timer         *Timer
	Counter       *Counter
	Process       *Process
}

var (
	wgSch   sync.WaitGroup
	jobsSch = make(chan string)
	Mu      sync.Mutex
)

//...
// Functions
func (s *Search) RunSearch() {
	s.initSearch()

	s.DrawSep("RECHERCHES")
	s.Timer.SearchStart = time.Now()

	s.loopDirsWorker(s.SrcPath)

	wgSch.Wait()

	fmt.Print("\r                                                                                                 ")

	s.Timer.SearchEnd = time.Since(s.Timer.SearchStart)

	time.Sleep(1 * time.Millisecond)

	s.DrawSep("EXPORT XLSX")
	if !s.Silent {
		RunWritter()
	} else {
		print("\n")
	}
}

func (s *Search) initSearch() {
	s.DrawSep("PARAMETRES")

	s.DrawParam("INITIALISATION DE LA RECHERCHE EN COURS")

	// Construct variable of search
	s.ReqUse = s.getReqOfSearched()
	if !s.Maj {
		s.Word = StrToLower(s.Word)
	}
	s.Ext = fmt.Sprintf(".%s", s.Ext)

	// Add WhiteList / BlackList
	if s.WordsList {
		s.setList(filepath.Join(s.DstPath, "words.txt"), 0)
		s.DrawParam(fmt.Sprintf("WORDS: %v", s.ListWords))
	}
	if s.BlackList {
		blPath := filepath.Join(s.DstPath, "blacklist")
		s.setList(filepath.Join(blPath, "__ALL__.txt"), 1)

		file := filepath.Join(blPath, fmt.Sprintf("%s.txt", StrToLower(s.Word)))
		if _, err := os.Stat(file); err == nil {
			s.setList(file, 1)
		}

		s.DrawParam(fmt.Sprintf("BLACKLIST: %v", s.ListBlackList))
	}
	if s.WhiteList {
		wlPath := filepath.Join(s.DstPath, "whitelist")
		s.setList(filepath.Join(wlPath, "__ALL__.txt"), 2)

		file := filepath.Join(wlPath, fmt.Sprintf("%s.txt", StrToLower(s.Word)))
		if _, err := os.Stat(file); err == nil {
			s.setList(file, 2)
		}

		s.DrawParam(fmt.Sprintf("WHITELIST: %v", s.ListWhiteList))
	}

	// Check basics configurations
	s.checkMinimumPoolSize()
	s.setMaxThread()

	// Creation of workers for search
	for w := 1; w <= s.PoolSize; w++ {
		numWorker := w
		go func() {
			err := s.loopFilesWorker()
			if err != nil {
				loger.Error(fmt.Sprintf("Error with worker N°%v", numWorker), err)
			}
		}()
	}

	// Create csv dump
	loger.Semicolon("id;Fichier;Date;Lien_Fichier;Lien")
}

func (s *Search) getReqOfSearched() string {

	req := "FilesDIR"

	if !s.Cls && !s.Compiler {
		req += fmt.Sprintf(" -mode=%s", s.Mode)

		if s.Word != "" {
			req += fmt.Sprintf(" -word=%s", s.Word)
		}

		if s.Ext != "" {
			req += fmt.Sprintf(" -ext=%s", s.Ext)
		}

		req += fmt.Sprintf(" -poolsize=%v", s.PoolSize)

		if s.Maj {
			req += " -maj"
		}

		if s.Devil {
			req += " -devil"
		}

		if s.BlackList {
			req += " -b"
		}

		if s.WhiteList {
			req += " -w"
		}
	}

	if s.Cls {
		req += " -cls"
	} else if s.Compiler {
		req += " -c"
	}

	if s.Silent {
		req += " -s"
	}

	s.DrawParam(fmt.Sprintf("REQUETE UTILISEE: %s", req))

	return req
}

func (s *Search) setList(file string, val int) {
	readFile, err := os.Open(file)
	if err != nil {
		list := "Words"
		if val == 1 {
			list = "Black"
		} else if val == 2 {
			list = "White"
		}
		loger.Error(fmt.Sprintf("Error during insert data in %sList:", list), err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		switch val {
		case 0:
			s.ListWords = append(s.ListWords, fileScanner.Text())
		case 1:
			s.ListWhiteList = append(s.ListWhiteList, fileScanner.Text())
		case 2:
			s.ListWhiteList = append(s.ListWhiteList, fileScanner.Text())
		}

	}
	_ = readFile.Close()
}

func (s *Search) checkMinimumPoolSize() {
	if s.PoolSize < 2 {
		s.PoolSize = 2
		s.DrawParam("POOLSIZE MISE A", strconv.Itoa(s.PoolSize), "(ne peut pas être inférieur)")
	} else {
		s.DrawParam("POOLSIZE MISE A", strconv.Itoa(s.PoolSize))
	}
}

func (s *Search) setMaxThread() {
	maxThr := s.PoolSize * 500
	debug.SetMaxThreads(maxThr)
	s.Process.NbrThreads = maxThr

	s.DrawParam("THREADS MIS A", strconv.Itoa(maxThr))
}

func (s *Search) isInWordsList(f string) bool {
	for _, word := range s.ListWords {

		if !s.Maj {
			f = StrToLower(f)
			word = StrToLower(word)
		}

		if strings.Contains(f, word) {
			return true
		}
	}
	return false
}

func (s *Search) isInBlackList(f string) bool {
	for _, word := range s.ListBlackList {
		if strings.Contains(StrToLower(f), StrToLower(word)) {
			return true
		}
	}
	return false
}

func (s *Search) isInWhiteList(f string) bool {
	for _, word := range s.ListWhiteList {
		if strings.Contains(StrToLower(f), StrToLower(word)) {
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

	//TODO: Tout refaire avec du REGEX

	// condition of search Mode ( = | % | ^ | $ )
	switch s.Mode {
	case "%":
		if s.WordsList {
			return s.isInWordsList(name)
		} else {
			if !strings.Contains(name, s.Word) {
				return false
			}
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
func (s *Search) loopFilesWorker() error {
	for jobPath := range jobsSch {

		if s.BlackList && s.isInBlackList(jobPath) {
			wgSch.Done()
			continue
		}
		if s.WhiteList && !s.isInWhiteList(jobPath) {
			wgSch.Done()
			continue
		}

		files, err := ioutil.ReadDir(jobPath)
		if err != nil {
			loger.Crash("Crash with this path:", err)
			wgSch.Done()
			return err
		}

		for _, file := range files {
			if !file.IsDir() {
				if s.checkFileSearched(file.Name()) {
					Mu.Lock()
					atomic.AddUint64(&s.Counter.NbrFiles, 1)

					if !s.Silent {
						s.DrawFilesOk(file.Name())
					}

					loger.Semicolon(fmt.Sprintf("%v;%s;%s;%s;%s", s.Counter.NbrFiles, file.Name(), file.ModTime().Format("02-01-2006 15:04:05"), filepath.Join(jobPath, file.Name()), jobPath))

					data := Export{
						Id:       int(s.Counter.NbrFiles),
						File:     file.Name(),
						Date:     file.ModTime().Format("02-01-2006 15:04:05"),
						PathFile: filepath.Join(jobPath, file.Name()),
						Path:     jobPath,
					}
					ExportSch = append(ExportSch, data)

					//  TODO: Ajouter les mode -S dans le drawings pour les prints
					Mu.Unlock()
				}
				atomic.AddUint64(&s.Counter.NbrAllFiles, 1)

				if runtime.NumGoroutine() > s.Process.NbrGoroutines {
					s.Process.NbrGoroutines = runtime.NumGoroutine()
				}
			} else {
				atomic.AddUint64(&s.Counter.NbrFolder, 1)
			}
			s.DrawFilesSearched()
		}
		wgSch.Done()
	}
	return nil
}

func (s *Search) loopDirsWorker(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		loger.Error("Error with this path:", err)
	}

	go func() {
		wgSch.Add(1)
		jobsSch <- path
	}()

	for _, file := range files {
		if file.IsDir() {
			if s.Devil {
				time.Sleep(20 * time.Millisecond)
				go s.loopDirsWorker(filepath.Join(path, file.Name()))
			} else {
				s.loopDirsWorker(filepath.Join(path, file.Name()))
			}
		}
	}
}
