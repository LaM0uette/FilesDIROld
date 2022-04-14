package pkg

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
)

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
	BlackList bool
	WhiteList bool

	// Search
	SrcPath string
	DstPath string
	ReqUse  string

	// Data
	ListBlackList []string
	ListWhiteList []string
}

//...
// Functions
func (s *Search) RunSearch() {
	s.initSearch()
}

func (s *Search) initSearch() {
	DrawParam("INITIALISATION DE LA RECHERCHE EN COURS")

	// Construct variable of search
	s.ReqUse = s.GetReqOfSearched()
	if !s.Maj {
		s.Word = StrToLower(s.Word)
	}
	s.Ext = fmt.Sprintf(".%s", s.Ext)

	// Add WhiteList / BlackList
	if s.BlackList {
		blPath := filepath.Join(s.DstPath, "blacklist")
		s.setBlackWhiteList(filepath.Join(blPath, "__ALL__.txt"), 0)

		file := filepath.Join(blPath, fmt.Sprintf("%s.txt", StrToLower(s.Word)))
		if _, err := os.Stat(file); err == nil {
			s.setBlackWhiteList(file, 0)
		}

		DrawParam(fmt.Sprintf("BLACKLIST: %v", s.ListBlackList))
	}
	if s.WhiteList {
		wlPath := filepath.Join(s.DstPath, "whitelist")
		s.setBlackWhiteList(filepath.Join(wlPath, "__ALL__.txt"), 1)

		file := filepath.Join(wlPath, fmt.Sprintf("%s.txt", StrToLower(s.Word)))
		if _, err := os.Stat(file); err == nil {
			s.setBlackWhiteList(file, 1)
		}

		DrawParam(fmt.Sprintf("WHITELIST: %v", s.ListWhiteList))
	}

	// Check basics configurations
	s.CheckMinimumPoolSize()
	s.SetMaxThread()
}

func (s *Search) GetReqOfSearched() string {

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

	DrawParam(fmt.Sprintf("REQUETE UTILISEE: %s", req))

	return req
}

func (s *Search) setBlackWhiteList(file string, val int) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err) //TODO: Loger crash
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		switch val {
		case 0:
			s.ListBlackList = append(s.ListBlackList, fileScanner.Text())
		case 1:
			s.ListWhiteList = append(s.ListWhiteList, fileScanner.Text())
		}

	}
	_ = readFile.Close()
}

func (s *Search) CheckMinimumPoolSize() {
	if s.PoolSize < 2 {
		s.PoolSize = 2
		DrawParam("POOLSIZE MISE A", strconv.Itoa(s.PoolSize), "(ne peut pas être inférieur)")
	} else {
		DrawParam("POOLSIZE MISE A", strconv.Itoa(s.PoolSize))
	}
}

func (s *Search) SetMaxThread() {
	maxThr := s.PoolSize * 500
	debug.SetMaxThreads(maxThr)

	if s.Silent {
		return
	}

	//loger.Paramln(display.DrawSetMaxThread(maxThr))
}
