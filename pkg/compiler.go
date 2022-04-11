package pkg

import (
	"FilesDIR/display"
	"FilesDIR/globals"
	"FilesDIR/loger"
	"FilesDIR/task"
	"fmt"
	"github.com/qax-os/excelize"
	"github.com/tealeg/xlsx"
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
	Wb   = &xlsx.File{}
	Sht  = &xlsx.Sheet{}
	Id   int
)

//...
// ACTIONS:
func ClsTempFiles() {
	_ = os.RemoveAll(globals.FolderLogs)
	_ = os.RemoveAll(globals.FolderDumps)
	_ = os.RemoveAll(globals.FolderExports)
}

func setCellValue(r, c int, val any) {
	cell, _ := Sht.Cell(r, c)
	cell.SetValue(val)
}

func getCellValue(sht *xlsx.Sheet, r, c int) (str string, style *xlsx.Style) {
	cell, _ := sht.Cell(r, c)
	return cell.Value, cell.GetStyle()
}

func setCellBgColor(r, c int, style *xlsx.Style) {
	cell, _ := Sht.Cell(r, c)
	cell.SetStyle(style)
}

func CompilerFicheAppuiFt(path string) {
	path = "C:\\Users\\XD5965\\OneDrive - EQUANS\\Bureau\\Nouveau dossier"

	loger.BlankDateln(display.DrawInitCompiler())
	time.Sleep(800 * time.Millisecond)

	loger.Blankln(display.DrawRunCompiler())

	Wb = xlsx.NewFile()
	Sht, _ = Wb.AddSheet("Sheet1")

	setCellValue(0, 0, "Chemin de la fiche")
	setCellValue(0, 1, "Adresse")
	setCellValue(0, 2, "Ville")
	setCellValue(0, 3, "Num appui")
	setCellValue(0, 4, "Type appui")
	setCellValue(0, 5, "Type_n_app")
	setCellValue(0, 6, "Nature TVX")
	setCellValue(0, 7, "Etiquette jaune")
	setCellValue(0, 8, "Effort avant ajout câble")
	setCellValue(0, 9, "Effort après ajout câble")
	setCellValue(0, 10, "Effort nouveau appui")
	setCellValue(0, 11, "Latitude")
	setCellValue(0, 12, "Longitude")
	setCellValue(0, 13, "Opérateur")
	setCellValue(0, 14, "Appui utilisable en l'état")
	setCellValue(0, 15, "Environnement")
	setCellValue(0, 16, "Commentaire appui")
	setCellValue(0, 17, "Commentaire global")
	setCellValue(0, 18, "Proxi ENEDIS")
	setCellValue(0, 19, "id_metier_")
	setCellValue(0, 20, "Date")
	setCellValue(0, 21, "PB")

	for w := 1; w <= 500; w++ {
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
				loger.Errorln(fmt.Sprintf("Crash with this files: %s", excelFile))
				continue
			}

			rows, err := f.GetRows("Sheet1")
			if err != nil {
				loger.Errorln(err)
				continue
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
	time.Sleep(500 * time.Millisecond)

	loger.BlankDateln(display.DrawEndCompiler())

	loger.BlankDateln(fmt.Sprintf("Nombre de fiches compilées : %v", Id-1))
	time.Sleep(800 * time.Millisecond)

	err = Wb.Save(filepath.Join(path, fmt.Sprintf("__COMPILATION__%v.xlsx", time.Now().Format("20060102150405"))))
	if err != nil {
		return
	}

	loger.Blankln(display.DrawSaveExcel())
	fmt.Println()
	time.Sleep(200 * time.Millisecond)
}

//...
//WORKER:
func workerFicheAppuiFt() {
	for job := range jobs {
		loger.BlankDateln(fmt.Sprintf("N°%v | Files: %s", job.Id, filepath.Base(job.Path)))

		excelFile := job.Path
		f, err := xlsx.OpenFile(excelFile)
		if err != nil {
			loger.Errorln(fmt.Sprintf("Crash with this files: %s", filepath.Base(excelFile)))
			wg.Done()
			continue
		}

		sht := f.Sheets[0]

		verifFA, _ := getCellValue(sht, 0, 0)
		if !strings.Contains(task.StrToLower(verifFA), "appui") {
			loger.Errorln(fmt.Sprintf("Ce n'est pas le bon format de fiche appui: %s", filepath.Base(excelFile)))
			wg.Done()
			continue
		}

		adresse, _ := getCellValue(sht, 4, 3)
		ville, _ := getCellValue(sht, 3, 3)
		numAppui, _ := getCellValue(sht, 2, 3)
		type1, _ := getCellValue(sht, 25, 2)
		typeNApp, _ := getCellValue(sht, 51, 12)
		natureTvx, _ := getCellValue(sht, 52, 12)

		etiquetteJaune, _ := getCellValue(sht, 11, 20)
		switch task.StrToLower(etiquetteJaune) {
		case "oui":
			etiquetteJaune = "non"
		case "non":
			etiquetteJaune = "oui"
		}

		effort1, rgb1 := getCellValue(sht, 25, 18)
		effort2, rgb2 := getCellValue(sht, 25, 20)
		effort3, rgb3 := getCellValue(sht, 25, 22)

		lat, _ := getCellValue(sht, 4, 15)
		lon, _ := getCellValue(sht, 5, 15)
		operateur, _ := getCellValue(sht, 2, 9)
		utilisableEnEtat, _ := getCellValue(sht, 11, 22)
		environnement, _ := getCellValue(sht, 51, 22)
		commentaireEtatAppui, _ := getCellValue(sht, 12, 5)
		commentaireGlobal, _ := getCellValue(sht, 54, 0)
		proxiEnedis, _ := getCellValue(sht, 52, 22)

		insee, _ := getCellValue(sht, 3, 21)
		idMetier := fmt.Sprintf("%s/%s", numAppui, insee)
		date, _ := getCellValue(sht, 0, 19)
		pb, _ := getCellValue(sht, 17, 13)

		// insert value
		setCellValue(job.Id, 0, job.Path)
		setCellValue(job.Id, 1, adresse)
		setCellValue(job.Id, 2, ville)
		setCellValue(job.Id, 3, numAppui)
		setCellValue(job.Id, 4, type1)
		setCellValue(job.Id, 5, typeNApp)
		setCellValue(job.Id, 6, natureTvx)
		setCellValue(job.Id, 7, etiquetteJaune)
		setCellValue(job.Id, 8, effort1)
		setCellValue(job.Id, 9, effort2)
		setCellValue(job.Id, 10, effort3)
		setCellValue(job.Id, 11, lat)
		setCellValue(job.Id, 12, lon)
		setCellValue(job.Id, 13, operateur)
		setCellValue(job.Id, 14, utilisableEnEtat)
		setCellValue(job.Id, 15, environnement)
		setCellValue(job.Id, 16, commentaireEtatAppui)
		setCellValue(job.Id, 17, commentaireGlobal)
		setCellValue(job.Id, 18, proxiEnedis)
		setCellValue(job.Id, 19, idMetier)
		setCellValue(job.Id, 20, date)
		setCellValue(job.Id, 21, pb)

		setCellBgColor(job.Id, 8, rgb1)
		setCellBgColor(job.Id, 9, rgb2)
		setCellBgColor(job.Id, 10, rgb3)

		wg.Done()
	}
}
