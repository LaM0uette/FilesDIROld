package pkg

import (
	"FilesDIR/display"
	"FilesDIR/globals"
	"FilesDIR/loger"
	"FilesDIR/task"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type compilData struct {
	Path string
	Id   int
}

var (
	wg   sync.WaitGroup
	jobs = make(chan compilData)
	Wb   = &excelize.File{}
	Id   int
)

//...
// ACTIONS:
func ClsTempFiles() {
	_ = os.RemoveAll(globals.FolderLogs)
	_ = os.RemoveAll(globals.FolderDumps)
	_ = os.RemoveAll(globals.FolderExports)
}

func CompilerFicheAppuiFt(path string) {

	loger.BlankDateln(display.DrawInitCompiler())
	time.Sleep(800 * time.Millisecond)

	loger.BlankDateln(display.DrawRunCompiler())

	Id = 1

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
		if !file.IsDir() && !strings.Contains(file.Name(), "__COMPILATION__") {

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
					Id++

					a := compilData{
						Path: r,
						Id:   Id,
					}

					jobs <- a
				}()
			}
		}
	}

	wg.Wait()
	time.Sleep(1 * time.Second)

	loger.BlankDateln(display.DrawEndCompiler())

	loger.BlankDateln(fmt.Sprintf("Nombre de lignes compilées : %v", Id-1))
	time.Sleep(800 * time.Millisecond)

	if err := Wb.SaveAs(filepath.Join(path, fmt.Sprintf("__COMPILATION__%v.xlsx", time.Now().Format("20060102150405")))); err != nil {
		fmt.Println(err)
	}

	loger.Blankln(display.DrawSaveExcel())
	fmt.Println()
	time.Sleep(200 * time.Millisecond)
}

//...
//WORKER:
func workerFicheAppuiFt() {
	for job := range jobs {
		loger.BlankDateln(fmt.Sprintf("\rN° de la lignes compilées :  %v", job.Id-1))

		excelFile := job.Path
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

		// insert value
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", job.Id), job.Path)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", job.Id), adresse)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", job.Id), ville)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", job.Id), numAppui)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", job.Id), type1)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("F%v", job.Id), typeNApp)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("G%v", job.Id), natureTvx)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("H%v", job.Id), etiquetteJaune)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("I%v", job.Id), effort1)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("J%v", job.Id), effort2)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("K%v", job.Id), effort3)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("L%v", job.Id), lat)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("M%v", job.Id), lon)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("N%v", job.Id), operateur)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("O%v", job.Id), utilisableEnEtat)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("P%v", job.Id), environnement)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("Q%v", job.Id), commentaireEtatAppui)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("R%v", job.Id), commentaireGlobal)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("S%v", job.Id), proxiEnedis)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("T%v", job.Id), idMetier)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("U%v", job.Id), date)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("V%v", job.Id), pb)

		wg.Done()
	}
}
