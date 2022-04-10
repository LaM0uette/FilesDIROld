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
)

var (
	wg   sync.WaitGroup
	jobs = make(chan int)
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
	//for w := 1; w <= 300; w++ {
	//	go workerFicheAppuiFt(Wb)
	//}

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
				fmt.Println(row[3])
			}
		}
	}

	//for i := 0; i < iMax; i++ {
	//	i := i
	//	go func() {
	//		wg.Add(1)
	//		jobs <- i
	//	}()
	//}

	//wg.Wait()
	//time.Sleep(1 * time.Second)
	//fmt.Printf("\rNombre de lignes compilées :  %v/%v\n", iMax, iMax)
}

//...
// WORKER:
//func workerFicheAppuiFt(Wb *excelize.File) {
//	for job := range jobs {
//
//		fmt.Print("\r")
//		fmt.Printf("\rCompilation du fichier Excel...  %v/%v", job, iMax)
//
//		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", job+2), ExcelData[job].Id)
//		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", job+2), ExcelData[job].File)
//		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", job+2), ExcelData[job].Date)
//		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", job+2), ExcelData[job].PathFile)
//		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", job+2), ExcelData[job].Path)
//
//		wg.Done()
//	}
//}
