package construct

import (
	"FilesDIR/display"
	"FilesDIR/loger"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"path/filepath"
	"runtime/debug"
	"sync"
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
	FlgCls       bool
	FlgCompiler  bool
}

type ExportData struct {
	Id       int    `json:"id"`
	File     string `json:"Fichier"`
	Date     string `json:"Date"`
	PathFile string `json:"Lien_Fichier"`
	Path     string `json:"Lien"`
}

var (
	wg   sync.WaitGroup
	jobs = make(chan int)
	Wb   = &excelize.File{}

	ExcelData []ExportData
)

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

func (f *Flags) CheckMinimumPoolSize() {
	if f.FlgPoolSize < 2 {
		f.FlgPoolSize = 2
		loger.Paramln(display.DrawCheckMinimumPoolSize())
	}
}

func (f *Flags) SetMaxThread() {
	maxThr := f.FlgPoolSize * 500
	debug.SetMaxThreads(maxThr)

	if f.FlgSuper {
		return
	}

	loger.Paramln(fmt.Sprintf("<fg=214,99,144>Nombre de threads mis à :</> <fg=44,168,65>%v</>", maxThr))
}

func (f *Flags) SetSaveWord() string {
	word := f.FlgWord
	if len(f.FlgWord) < 1 {
		word = "Export"
		time.Sleep(600 * time.Millisecond)
		loger.Actionln(fmt.Sprintf("<fg=214,99,144>Nom du fichier de sauvergarde mis par défaut :</> <fg=44,168,65>%v</>", word))
	}

	return word
}

func (f *Flags) GenerateExcelSave(DstPath string) {
	if f.FlgXl || f.FlgSuper {
		return
	}

	Wb = excelize.NewFile()
	_ = Wb.SetCellValue("Sheet1", "A1", "id")
	_ = Wb.SetCellValue("Sheet1", "B1", "Fichier")
	_ = Wb.SetCellValue("Sheet1", "C1", "Date")
	_ = Wb.SetCellValue("Sheet1", "D1", "LienFichier")
	_ = Wb.SetCellValue("Sheet1", "E1", "Lien")

	f.DrawWriteExcel()

	// Creation of workers for write line in excel file
	iMax := len(ExcelData)
	for w := 1; w <= 300; w++ {
		go f.writeExcelLineWorker(Wb, iMax)
	}
	// Run writing loop
	for i := 0; i < iMax; i++ {
		i := i
		go func() {
			wg.Add(1)
			jobs <- i
		}()
	}

	wg.Wait() // Wait for all write loops to complete
	time.Sleep(300 * time.Millisecond)
	loger.POAction(fmt.Sprintf("\r<fg=214,99,144>Nombre de lignes sauvegardées :</>  <fg=44,168,65>%v</><fg=214,99,144>/</><fg=44,168,65>%v</>\n", iMax, iMax))
	time.Sleep(1 * time.Second)

	// Generate a default word if is none
	saveWord := f.SetSaveWord()

	// Save Excel file
	if err := Wb.SaveAs(filepath.Join(DstPath, saveWord+fmt.Sprintf("_%v.xlsx", time.Now().Format("20060102150405")))); err != nil {
		loger.Errorln(err)
	}

	f.DrawSaveExcel()
}

//...
// WORKER:
func (f *Flags) writeExcelLineWorker(Wb *excelize.File, iMax int) {
	for job := range jobs {

		//fmt.Print("\r")
		loger.POAction(fmt.Sprintf("\r<fg=214,99,144>Export Excel...</>  <fg=44,168,65>%v</><fg=214,99,144>/</><fg=44,168,65>%v</>", job, iMax))

		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", job+2), ExcelData[job].Id)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", job+2), ExcelData[job].File)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", job+2), ExcelData[job].Date)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", job+2), ExcelData[job].PathFile)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", job+2), ExcelData[job].Path)

		wg.Done()
	}
}

//...
// DRAWS:
func (f *Flags) DrawStart() {
	if f.FlgSuper {
		return
	}
	loger.Start(display.DrawStart())
	time.Sleep(1 * time.Second)
}

func (f *Flags) DrawInitSearch() {
	if f.FlgSuper {
		return
	}

	loger.Paramln(display.DrawInitSearch())
	time.Sleep(800 * time.Millisecond)
}

func (f *Flags) DrawRunSearch() {
	if f.FlgSuper {
		return
	}

	loger.Uiln(display.DrawRunSearch())
	time.Sleep(400 * time.Millisecond)
}

func (f *Flags) DrawEndSearch() {
	if f.FlgSuper {
		return
	}

	time.Sleep(1 * time.Second)
	loger.Uiln(display.DrawEndSearch())
	time.Sleep(200 * time.Millisecond)
}

func (f *Flags) DrawWriteExcel() {
	if f.FlgSuper {
		return
	}

	loger.Action(display.DrawWriteExcel())
	time.Sleep(200 * time.Millisecond)
}

func (f *Flags) DrawSaveExcel() {
	if f.FlgSuper {
		return
	}
	//fmt.Println()
	loger.Actionln(display.DrawSaveExcel())
	time.Sleep(200 * time.Millisecond)
}

func (f *Flags) DrawEnd(SrcPath, DstPath, ReqFinal string, NbGoroutine, NbFiles int, TimerSearch, timerEnd time.Duration) {
	disp := display.DrawEnd(SrcPath, DstPath, ReqFinal, NbGoroutine, NbFiles, f.FlgPoolSize, TimerSearch, timerEnd)
	loger.Endln(disp)
}

//...
// Pkg
func DrawEndCls() {
	loger.Actionln("*** Dossiers de logs et dumps nettoyés ! ***\n")
}
