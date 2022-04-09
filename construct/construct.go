package construct

import (
	"FilesDIR/display"
	"FilesDIR/log"
	"FilesDIR/loger"
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
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

//...
// ACTIONS:
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

func (f *Flags) ExportExcelActivate() bool {
	if f.FlgXl && f.FlgSuper {
		return false
	}
	return true
}

func (f *Flags) CheckMinimumPoolSize() {
	if f.FlgPoolSize < 2 {
		f.FlgPoolSize = 2
		loger.Info("Set the PoolSize to 2")
	}
}

func (f *Flags) SetMaxThread() {
	maxThr := f.FlgPoolSize * 500
	debug.SetMaxThreads(maxThr)

	if f.FlgSuper {
		return
	}

	loger.Info(fmt.Sprintf("Set max thread count to %v", maxThr))
}

//...
// DRAWS:
func (f *Flags) DrawStart() {
	if f.FlgSuper {
		return
	}
	loger.Blank(display.DrawStart())
	time.Sleep(1 * time.Second)
}

func (f *Flags) DrawInitSearch() {
	if f.FlgSuper {
		return
	}

	loger.BlankDate(display.DrawInitSearch())
	time.Sleep(800 * time.Millisecond)
}

func (f *Flags) DrawRunSearch() {
	if f.FlgSuper {
		return
	}

	loger.Blank(display.DrawRunSearch())
	time.Sleep(400 * time.Millisecond)
}

func (f *Flags) DrawEndSearch() {
	if f.FlgSuper {
		return
	}

	time.Sleep(1 * time.Second)
	loger.Blank(display.DrawEndSearch())
	time.Sleep(200 * time.Millisecond)
}

func (f *Flags) DrawEnd(SrcPath, DstPath, ReqFinal string, NbGoroutine, NbFiles int, TimerSearch, timerEnd time.Duration) {
	disp := display.DrawEnd(SrcPath, DstPath, ReqFinal, NbGoroutine, NbFiles, f.FlgPoolSize, TimerSearch, timerEnd)
	log.Blank.Print(disp)
	fmt.Print(disp)

	fmt.Print("Appuyer sur Entrée pour quitter...")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		log.Crash.Println(err)
	}
}
