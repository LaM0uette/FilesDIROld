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
	Super     bool
	BlackList bool

	// Search
	SrcPath string
	DstPath string
	ReqUse  string

	// Data
	ListBlackList []string
}

//...
// Functions
func (s *Search) RunSearch() {
	s.initSearch()
}

func (s *Search) initSearch() {
	DrawParam("INITIALISATION DE LA RECHERCHE EN COURS")

	s.ReqUse = s.GetReqOfSearched()
	if !s.Maj {
		s.Word = StrToLower(s.Word)
	}
	s.Ext = fmt.Sprintf(".%s", s.Ext)

	if s.BlackList {
		blPath := filepath.Join(s.DstPath, "blacklist")
		s.setBlackList(filepath.Join(blPath, "__ALL__.txt"))

		file := filepath.Join(blPath, fmt.Sprintf("%s.txt", StrToLower(s.Word)))
		if _, err := os.Stat(file); err == nil {
			s.setBlackList(file)
		}
	}

	s.CheckMinimumPoolSize()
	s.SetMaxThread()
}

func (s *Search) GetReqOfSearched() string {
	DrawParam("CONSTRUCTION DE LA REQUETE EN COURS")

	VWord := ""
	if s.Word != "" {
		VWord = " -word=" + s.Word
	}

	VMaj := ""
	if s.Maj {
		VMaj = " -maj"
	}

	VDevil := ""
	if s.Devil {
		VDevil = " -devil"
	}

	VSuper := ""
	if s.Super {
		VSuper = " -s"
	}

	VBlackList := ""
	if s.BlackList {
		VBlackList = " -b"
	}
	return fmt.Sprintf("FilesDIR -mode=%s%s -ext=%s -poolsize=%v%s%s%s%s\n", s.Mode, VWord, s.Ext, s.PoolSize, VMaj, VDevil, VSuper, VBlackList)
}

func (s *Search) setBlackList(file string) {
	DrawParam("AJOUT DE DONNEES DANS LA BLACKLIST")

	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err) //TODO: Loger crash
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s.ListBlackList = append(s.ListBlackList, fileScanner.Text())
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

	if s.Super {
		return
	}

	//loger.Paramln(display.DrawSetMaxThread(maxThr))
}
