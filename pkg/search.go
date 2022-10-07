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

var (
	wgSch   sync.WaitGroup
	jobsSch = make(chan string)
	Mu      sync.Mutex
)

//...
// Functions
func (flg *SSearch) RunSearch() {
	flg.initSearch()

	flg.DrawSep("RECHERCHES")
	flg.Timer.SearchStart = time.Now()

	flg.loopDirsWorker(flg.SrcPath)

	wgSch.Wait()

	fmt.Print("\r                                                                                                 ")

	flg.Timer.SearchEnd = time.Since(flg.Timer.SearchStart)

	time.Sleep(1 * time.Millisecond)

	flg.DrawSep("EXPORT XLSX")
	if !flg.Silent {
		RunWritter()
	} else {
		print("\n")
	}
}

func (flg *SSearch) initSearch() {
	flg.DrawSep("PARAMETRES")

	flg.DrawParam("INITIALISATION DE LA RECHERCHE EN COURS")

	// Construct variable of search
	flg.ReqUse = flg.getReqOfSearched()
	if !flg.Maj {
		flg.Word = StrToLower(flg.Word)
	}
	flg.Ext = fmt.Sprintf(".%s", flg.Ext)

	// Add WhiteList / BlackList
	if flg.WordsList {
		flg.setList(filepath.Join(flg.DstPath, "words.txt"), 0)
		flg.DrawParam(fmt.Sprintf("WORDS: %v", flg.ListWords))
	}
	if flg.BlackList {
		blPath := filepath.Join(flg.DstPath, "blacklist")
		flg.setList(filepath.Join(blPath, "__ALL__.txt"), 1)

		file := filepath.Join(blPath, fmt.Sprintf("%s.txt", StrToLower(flg.Word)))
		if _, err := os.Stat(file); err == nil {
			flg.setList(file, 1)
		}

		flg.DrawParam(fmt.Sprintf("BLACKLIST: %v", flg.ListBlackList))
	}
	if flg.WhiteList {
		wlPath := filepath.Join(flg.DstPath, "whitelist")
		flg.setList(filepath.Join(wlPath, "__ALL__.txt"), 2)

		file := filepath.Join(wlPath, fmt.Sprintf("%s.txt", StrToLower(flg.Word)))
		if _, err := os.Stat(file); err == nil {
			flg.setList(file, 2)
		}

		flg.DrawParam(fmt.Sprintf("WHITELIST: %v", flg.ListWhiteList))
	}

	// Check basics configurations
	flg.checkMinimumPoolSize()
	flg.setMaxThread()

	// Creation of workers for search
	for w := 1; w <= flg.PoolSize; w++ {
		numWorker := w
		go func() {
			err := flg.loopFilesWorker()
			if err != nil {
				loger.Errorinf(fmt.Sprintf("Error with worker N°%v", numWorker), err)
			}
		}()
	}

	// Create csv dump
	loger.Semicolon("id;Fichier;Date;Lien_Fichier;Lien")
}

func (flg *SSearch) getReqOfSearched() string {

	req := "FilesDIR"

	if !flg.Cls && !flg.Compiler {
		req += fmt.Sprintf(" -mode=%s", flg.Mode)

		if flg.Word != "" {
			req += fmt.Sprintf(" -word=%s", flg.Word)
		}

		if flg.Ext != "" {
			req += fmt.Sprintf(" -ext=%s", flg.Ext)
		}

		req += fmt.Sprintf(" -poolsize=%v", flg.PoolSize)

		if flg.Maj {
			req += " -maj"
		}

		if flg.Devil {
			req += " -devil"
		}

		if flg.BlackList {
			req += " -b"
		}

		if flg.WhiteList {
			req += " -w"
		}
	}

	if flg.Cls {
		req += " -cls"
	} else if flg.Compiler {
		req += " -c"
	}

	if flg.Silent {
		req += " -s"
	}

	flg.DrawParam(fmt.Sprintf("REQUETE UTILISEE: %s", req))

	return req
}

func (flg *SSearch) setList(file string, val int) {
	readFile, err := os.Open(file)
	if err != nil {
		list := "Words"
		if val == 1 {
			list = "Black"
		} else if val == 2 {
			list = "White"
		}
		loger.Errorinf(fmt.Sprintf("Error during insert data in %sList:", list), err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		switch val {
		case 0:
			flg.ListWords = append(flg.ListWords, fileScanner.Text())
		case 1:
			flg.ListWhiteList = append(flg.ListWhiteList, fileScanner.Text())
		case 2:
			flg.ListWhiteList = append(flg.ListWhiteList, fileScanner.Text())
		}

	}
	_ = readFile.Close()
}

func (flg *SSearch) checkMinimumPoolSize() {
	if flg.PoolSize < 2 {
		flg.PoolSize = 2
		flg.DrawParam("POOLSIZE MISE A", strconv.Itoa(flg.PoolSize), "(ne peut pas être inférieur)")
	} else {
		flg.DrawParam("POOLSIZE MISE A", strconv.Itoa(flg.PoolSize))
	}
}

func (flg *SSearch) setMaxThread() {
	maxThr := flg.PoolSize * 500
	debug.SetMaxThreads(maxThr)
	flg.Process.NbrThreads = maxThr

	flg.DrawParam("THREADS MIS A", strconv.Itoa(maxThr))
}

func (flg *SSearch) isInWordsList(f string) bool {
	for _, word := range flg.ListWords {

		if !flg.Maj {
			f = StrToLower(f)
			word = StrToLower(word)
		}

		if strings.Contains(f, word) {
			return true
		}
	}
	return false
}

func (flg *SSearch) isInBlackList(f string) bool {
	for _, word := range flg.ListBlackList {
		if strings.Contains(StrToLower(f), StrToLower(word)) {
			return true
		}
	}
	return false
}

func (flg *SSearch) isInWhiteList(f string) bool {
	for _, word := range flg.ListWhiteList {
		if strings.Contains(StrToLower(f), StrToLower(word)) {
			return true
		}
	}
	return false
}

func (flg *SSearch) checkFileSearched(file string) bool {
	name := file[:strings.LastIndex(file, path.Ext(file))]
	ext := StrToLower(filepath.Ext(file))

	if !flg.Maj {
		name = StrToLower(name)
	}

	//TODO: Tout refaire avec du REGEX

	// condition of search Mode ( = | % | ^ | $ )
	switch flg.Mode {
	case "%":
		if flg.WordsList {
			return flg.isInWordsList(name)
		} else {
			if !strings.Contains(name, flg.Word) {
				return false
			}
		}
	case "=":
		if name != flg.Word {
			return false
		}
	case "^":
		if !strings.HasPrefix(name, flg.Word) {
			return false
		}
	case "$":
		if !strings.HasSuffix(name, flg.Word) {
			return false
		}
	default:
		if !strings.Contains(name, flg.Word) {
			return false
		}
	}

	// condition of extension file
	if flg.Ext != ".*" && ext != flg.Ext {
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
func (flg *SSearch) loopFilesWorker() error {
	for jobPath := range jobsSch {

		if flg.BlackList && flg.isInBlackList(jobPath) {
			wgSch.Done()
			continue
		}
		if flg.WhiteList && !flg.isInWhiteList(jobPath) {
			wgSch.Done()
			continue
		}

		files, err := ioutil.ReadDir(jobPath)
		if err != nil {
			loger.Errorinf("Error with this path:", err)
			wgSch.Done()
			return err
		}

		for _, file := range files {
			if !file.IsDir() {
				if flg.checkFileSearched(file.Name()) {
					Mu.Lock()
					atomic.AddUint64(&flg.Counter.NbrFiles, 1)

					if !flg.Silent {
						flg.DrawFilesOk(file.Name())
					}

					loger.Semicolon(fmt.Sprintf("%v;%s;%s;%s;%s", flg.Counter.NbrFiles, file.Name(), file.ModTime().Format("02-01-2006 15:04:05"), filepath.Join(jobPath, file.Name()), jobPath))

					data := Export{
						Id:       int(flg.Counter.NbrFiles),
						File:     file.Name(),
						Date:     file.ModTime().Format("02-01-2006 15:04:05"),
						PathFile: filepath.Join(jobPath, file.Name()),
						Path:     jobPath,
					}
					ExportSch = append(ExportSch, data)

					//  TODO: Ajouter les mode -S dans le drawings pour les prints
					Mu.Unlock()
				}
				atomic.AddUint64(&flg.Counter.NbrAllFiles, 1)

				if runtime.NumGoroutine() > flg.Process.NbrGoroutines {
					flg.Process.NbrGoroutines = runtime.NumGoroutine()
				}
			} else {
				atomic.AddUint64(&flg.Counter.NbrFolder, 1)
			}
			flg.DrawFilesSearched()
		}
		wgSch.Done()
	}
	return nil
}

func (flg *SSearch) loopDirsWorker(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		loger.Errorinf("Error with this path:", err)
	}

	go func() {
		wgSch.Add(1)
		jobsSch <- path
	}()

	for _, file := range files {
		if file.IsDir() {
			if flg.Devil {
				time.Sleep(20 * time.Millisecond)
				go flg.loopDirsWorker(filepath.Join(path, file.Name()))
			} else {
				flg.loopDirsWorker(filepath.Join(path, file.Name()))
			}
		}
	}
}
