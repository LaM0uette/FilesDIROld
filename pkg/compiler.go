package pkg

import (
	"FilesDIR/globals"
	"FilesDIR/loger"
	"FilesDIR/task"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	jobs = make(chan string)
	Wb   = &excelize.File{}
)

//...
// ACTIONS:
func ClsTempFiles() {
	_ = os.RemoveAll(globals.FolderLogs)
	_ = os.RemoveAll(globals.FolderDumps)
	_ = os.RemoveAll(globals.FolderExports)
}

func CompilerFicheAppuiFt(path string) {
	for w := 1; w <= 300; w++ {
		go workerFicheAppuiFt()
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		loger.Crashln(fmt.Sprintf("Crash with this path: %s", path))
	}

	for _, file := range files {
		if !file.IsDir() {

			excelFile := filepath.Join(path, file.Name())
			f, err := excelize.OpenFile(excelFile)
			if err != nil {
				loger.Crashln(fmt.Sprintf("Crash with this files: %s", excelFile))
			}

			rows, err := f.GetRows("Sheet1")
			if err != nil {
				loger.Crashln(err)
			}

			for ir, row := range rows {
				if ir == 0 {
					continue
				}

				r := row[3]
				go func() {
					wg.Add(1)
					jobs <- r
				}()
			}
		}
	}

	wg.Wait()
	time.Sleep(1 * time.Second)
	//fmt.Printf("\rNombre de lignes compilées :  %v/%v\n", iMax, iMax)
}

//...
//WORKER:
func workerFicheAppuiFt() {
	for job := range jobs {

		excelFile := job
		f, err := excelize.OpenFile(excelFile)
		if err != nil {
			loger.Crashln(fmt.Sprintf("Crash with this files: %s", excelFile))
		}

		sht := f.GetSheetName(f.GetActiveSheetIndex())

		adresse, _ := f.GetCellValue(sht, "D5")
		ville, _ := f.GetCellValue(sht, "D4")
		numAppui, _ := f.GetCellValue(sht, "D3")
		type1, _ := f.GetCellValue(sht, "C26")
		typeNApp, _ := f.GetCellValue(sht, "M52")
		natureTvx, _ := f.GetCellValue(sht, "M53")

		etiquetteJaune, _ := f.GetCellValue(sht, "U12")
		switch task.StrToLower(etiquetteJaune) {
		case "oui":
			etiquetteJaune = "non"
		case "non":
			etiquetteJaune = "oui"
		}

		effort1, _ := f.GetCellValue(sht, "S26")
		effort2, _ := f.GetCellValue(sht, "U26")
		effort3, _ := f.GetCellValue(sht, "W26")
		lat, _ := f.GetCellValue(sht, "P5")
		lon, _ := f.GetCellValue(sht, "P6")
		operateur, _ := f.GetCellValue(sht, "J3")
		utilisableEnEtat, _ := f.GetCellValue(sht, "W12")
		environnement, _ := f.GetCellValue(sht, "W52")
		commentaireEtatAppui, _ := f.GetCellValue(sht, "F13")
		commentaireGlobal, _ := f.GetCellValue(sht, "A55")
		proxiEnedis, _ := f.GetCellValue(sht, "W53")
		idMetier, _ := fmt.Sprintf("%s/%s", numAppui, f.GetCellValue(sht, "V4"))
		date, _ := f.GetCellValue(sht, "T1")
		pb, _ := f.GetCellValue(sht, "N18")

		wg.Done()
	}
}
