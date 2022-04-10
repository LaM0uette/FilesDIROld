package pkg

import (
	"FilesDIR/globals"
	"FilesDIR/loger"
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

		adresse, _ := Wb.GetCellValue(sht, "D5")

		fmt.Print("\r", adresse)

		wg.Done()
	}
}
