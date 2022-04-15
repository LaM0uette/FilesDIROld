package pkg

import (
	"FilesDIR/config"
	"FilesDIR/loger"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"path/filepath"
	"sync"
	"time"
)

type Export struct {
	Id       int    `json:"id"`
	File     string `json:"Fichier"`
	Date     string `json:"Date"`
	PathFile string `json:"Lien_Fichier"`
	Path     string `json:"Lien"`
}

var (
	ExportSch []Export
	Wb        *excelize.File

	wgWrt   sync.WaitGroup
	jobsWrt = make(chan int)
)

func RunWritter() {
	headers := map[string]string{
		"A1": "id",
		"B1": "Fichier",
		"C1": "Date",
		"D1": "LienFichier",
		"E1": "Lien",
	}

	Wb = excelize.NewFile()
	for header := range headers {
		_ = Wb.SetCellValue("Sheet1", header, headers[header])
	}

	// Creation of workers for write line in excel file
	for w := 1; w <= 300; w++ {
		go writeExcelWorker()
	}
	// Run writing loop
	for i := 0; i < len(ExportSch); i++ {
		i := i
		go func() {
			wgWrt.Add(1)
			jobsWrt <- i
		}()
	}

	wgWrt.Wait()
	time.Sleep(1 * time.Second)

	if err := Wb.SaveAs(filepath.Join(config.DstPath, "exports", fmt.Sprintf("Export_%v.xlsx", time.Now().Format("20060102150405")))); err != nil {
		loger.Error("Erreur pendant la sauvegarde du fichier Excel:", err)
	}

}

//...
// WORKER:
func writeExcelWorker() {
	for jobWrt := range jobsWrt {

		DrawAddLine(jobWrt, len(ExportSch))

		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", jobWrt+2), ExportSch[jobWrt].Id)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", jobWrt+2), ExportSch[jobWrt].File)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", jobWrt+2), ExportSch[jobWrt].Date)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", jobWrt+2), ExportSch[jobWrt].PathFile)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", jobWrt+2), ExportSch[jobWrt].Path)

		wgWrt.Done()
	}
}
