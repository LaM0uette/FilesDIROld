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
	wg      sync.WaitGroup
	jobs    = make(chan string)
	Wb      = &excelize.File{}
	AppuiId int
)

//...
// ACTIONS:
func ClsTempFiles() {
	_ = os.RemoveAll(globals.FolderLogs)
	_ = os.RemoveAll(globals.FolderDumps)
	_ = os.RemoveAll(globals.FolderExports)
}

func CompilerFicheAppuiFt(path string) {
	AppuiId = 1

	Wb = excelize.NewFile()
	_ = Wb.SetCellValue("Sheet1", "A1", "Chemin de la fiche")
	_ = Wb.SetCellValue("Sheet1", "B1", "Adresse")
	_ = Wb.SetCellValue("Sheet1", "C1", "Ville")
	_ = Wb.SetCellValue("Sheet1", "D1", "Num appui")
	_ = Wb.SetCellValue("Sheet1", "E1", "Type appui")
	_ = Wb.SetCellValue("Sheet1", "F1", "Type_n_app")
	_ = Wb.SetCellValue("Sheet1", "G1", "Nature TVX")
	_ = Wb.SetCellValue("Sheet1", "H1", "Etiquette jaune")
	_ = Wb.SetCellValue("Sheet1", "I1", "Effort avant ajout câble")
	_ = Wb.SetCellValue("Sheet1", "J1", "Effort après ajout câble")
	_ = Wb.SetCellValue("Sheet1", "K1", "Effort nouveau appui")
	_ = Wb.SetCellValue("Sheet1", "L1", "Latitude")
	_ = Wb.SetCellValue("Sheet1", "M1", "Longitude")
	_ = Wb.SetCellValue("Sheet1", "N1", "Opérateur")
	_ = Wb.SetCellValue("Sheet1", "O1", "Appui utilisable en l'état")
	_ = Wb.SetCellValue("Sheet1", "P1", "Environnement")
	_ = Wb.SetCellValue("Sheet1", "Q1", "Commentaire appui")
	_ = Wb.SetCellValue("Sheet1", "R1", "Commentaire global")
	_ = Wb.SetCellValue("Sheet1", "S1", "Proxi ENEDIS")
	_ = Wb.SetCellValue("Sheet1", "T1", "id_metier_")
	_ = Wb.SetCellValue("Sheet1", "U1", "Date")
	_ = Wb.SetCellValue("Sheet1", "V1", "PB")

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
	if err := Wb.SaveAs(filepath.Join(path, fmt.Sprintf("__COMPILATION__%v.xlsx", time.Now().Format("20060102150405")))); err != nil {
		fmt.Println(err)
	}
}

//...
//WORKER:
func workerFicheAppuiFt() {
	for job := range jobs {
		AppuiId++

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

		insee, _ := f.GetCellValue(sht, "V4")
		idMetier := fmt.Sprintf("%s/%s", numAppui, insee)

		date, _ := f.GetCellValue(sht, "T1")
		pb, _ := f.GetCellValue(sht, "N18")

		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", AppuiId), job)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", AppuiId), adresse)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", AppuiId), ville)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", AppuiId), numAppui)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", AppuiId), type1)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("F%v", AppuiId), typeNApp)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("G%v", AppuiId), natureTvx)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("H%v", AppuiId), etiquetteJaune)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("I%v", AppuiId), effort1)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("J%v", AppuiId), effort2)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("K%v", AppuiId), effort3)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("L%v", AppuiId), lat)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("M%v", AppuiId), lon)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("N%v", AppuiId), operateur)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("O%v", AppuiId), utilisableEnEtat)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("P%v", AppuiId), environnement)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("Q%v", AppuiId), commentaireEtatAppui)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("R%v", AppuiId), commentaireGlobal)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("S%v", AppuiId), proxiEnedis)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("T%v", AppuiId), idMetier)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("U%v", AppuiId), date)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("V%v", AppuiId), pb)

		wg.Done()
	}
}
