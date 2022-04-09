package construct

import (
	"FilesDIR/display"
	"FilesDIR/loger"
	"fmt"
	"time"
)

type Flags struct {
	FlgMode      string
	FlgWord      string
	FlgExt       string
	FlgPoolSize  int
	FlgPath      string
	FlgMaj       bool
	FlgXl        bool
	FlgDevil     bool
	FlgSuper     bool
	FlgBlackList bool
}

func (f *Flags) GetReqOfSearched() string {

	VWord := ""
	if f.FlgWord != "" {
		VWord = " -word=" + f.FlgWord
	}

	VMaj := ""
	if f.FlgMaj {
		VMaj = " -maj"
	}

	VXl := ""
	if f.FlgXl {
		VXl = " -xl"
	}

	VDevil := ""
	if f.FlgDevil {
		VDevil = " -devil"
	}

	VSuper := ""
	if f.FlgSuper {
		VSuper = " -s"
	}

	VBlackList := ""
	if f.FlgBlackList {
		VBlackList = " -b"
	}
	return fmt.Sprintf("FilesDIR -mode=%s%s -ext=%s -poolsize=%v%s%s%s%s%s\n", f.FlgMode, VWord, f.FlgExt, f.FlgPoolSize, VMaj, VXl, VDevil, VSuper, VBlackList)
}

func (f *Flags) DrawStart() {
	if f.FlgSuper {
		return
	}
	loger.Blank(display.DrawStart())
	time.Sleep(1 * time.Second)
}
