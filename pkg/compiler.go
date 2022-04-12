package pkg

import (
	"FilesDIR/display"
	"FilesDIR/globals"
	"FilesDIR/loger"
	"FilesDIR/task"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
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
	Wb   = &excelize.File{}
	Mu   sync.Mutex
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

	loger.Paramln(display.DrawInitCompiler())
	time.Sleep(800 * time.Millisecond)

	timeStart := time.Now()

	loger.Uiln(display.DrawRunCompiler())

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

	for w := 1; w <= 400; w++ {
		go workerFicheAppuiFt()
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		loger.Crashln(fmt.Sprintf("Crashln with this path: %s", path))
	}

	for _, file := range files {
		if !file.IsDir() && !strings.Contains(file.Name(), "__COMPILATION__") {

			excelFile := filepath.Join(path, file.Name())
			f, err := xlsx.OpenFile(excelFile)
			if err != nil {
				loger.Errorln(fmt.Sprintf("Crashln with this files: %s", excelFile))
				continue
			}

			sht := f.Sheets[0]

			maxRow := sht.MaxRow

			for i := 0; i < maxRow; i++ {
				row, err := sht.Row(i)
				if err != nil {
					panic(err)
				}

				if i == 0 {
					continue
				}

				go func() {
					Mu.Lock()
					wg.Add(1)
					Id++

					jobs <- compilData{
						Path: row.GetCell(3).String(),
						Id:   Id,
					}
					Mu.Unlock()
				}()
			}
		}
	}

	wg.Wait()

	timeEnd := time.Since(timeStart)

	time.Sleep(1 * time.Second)

	loger.Uiln(display.DrawEndCompiler())

	loger.Action(fmt.Sprintf("Nombre de fiches compilées : %v", Id-1))
	time.Sleep(800 * time.Millisecond)
	loger.Action(fmt.Sprintf("Temps écoulé : %v", timeEnd))
	time.Sleep(800 * time.Millisecond)

	if err := Wb.SaveAs(filepath.Join(path, fmt.Sprintf("__COMPILATION__%v.xlsx", time.Now().Format("20060102150405")))); err != nil {
		fmt.Println(err)
	}

	loger.Uiln(display.DrawSaveExcel())
	fmt.Println()
	time.Sleep(200 * time.Millisecond)
}

//...
//WORKER:
func workerFicheAppuiFt() {
	for job := range jobs {
		loger.Okln(fmt.Sprintf("N°%v | Files: %s", job.Id, filepath.Base(job.Path)))

		excelFile := job.Path
		f, err := xlsx.OpenFile(excelFile)
		if err != nil {
			loger.Errorln(fmt.Sprintf("Crashln with this files: %s", filepath.Base(excelFile)))
			_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", job.Id), job.Path)
			wg.Done()
			continue
		}

		sht := f.Sheets[0]

		adresse, _ := sht.Cell(4, 3)
		ville, _ := sht.Cell(3, 3)
		numAppui, _ := sht.Cell(2, 3)
		type1, _ := sht.Cell(25, 2)
		typeNApp, _ := sht.Cell(51, 12)
		natureTvx, _ := sht.Cell(52, 12)

		cellEtiquetteJaune, _ := sht.Cell(11, 20)
		etiquetteJaune := ""
		switch task.StrToLower(cellEtiquetteJaune.Value) {
		case "oui":
			etiquetteJaune = "non"
		case "non":
			etiquetteJaune = "oui"
		}

		effort1, _ := sht.Cell(25, 18)
		effort2, _ := sht.Cell(25, 20)
		effort3, _ := sht.Cell(25, 22)

		lat, _ := sht.Cell(4, 15)
		lon, _ := sht.Cell(5, 15)
		operateur, _ := sht.Cell(2, 9)
		utilisableEnEtat, _ := sht.Cell(11, 22)
		environnement, _ := sht.Cell(51, 22)
		commentaireEtatAppui, _ := sht.Cell(12, 5)
		commentaireGlobal, _ := sht.Cell(54, 0)
		proxiEnedis, _ := sht.Cell(52, 22)

		insee, _ := sht.Cell(3, 21)
		idMetier := fmt.Sprintf("%s/%s", numAppui.Value, insee.Value)
		date, _ := sht.Cell(0, 19)
		pb, _ := sht.Cell(17, 13)

		// insert value
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", job.Id), job.Path)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", job.Id), adresse.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", job.Id), ville.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", job.Id), numAppui.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", job.Id), type1.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("F%v", job.Id), typeNApp.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("G%v", job.Id), natureTvx.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("H%v", job.Id), etiquetteJaune)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("I%v", job.Id), effort1.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("J%v", job.Id), effort2.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("K%v", job.Id), effort3.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("L%v", job.Id), lat.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("M%v", job.Id), lon.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("N%v", job.Id), operateur.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("O%v", job.Id), utilisableEnEtat.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("P%v", job.Id), environnement.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("Q%v", job.Id), commentaireEtatAppui.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("R%v", job.Id), commentaireGlobal.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("S%v", job.Id), proxiEnedis.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("T%v", job.Id), idMetier)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("U%v", job.Id), date.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("V%v", job.Id), pb.Value)

		rgb1 := effort1.GetStyle().Fill.FgColor
		rgb2 := effort2.GetStyle().Fill.FgColor
		rgb3 := effort3.GetStyle().Fill.FgColor

		if len(rgb1) > 2 {
			style1, _ := Wb.NewStyle(fmt.Sprintf("{\"fill\":{\"type\":\"pattern\",\"color\":[\"#%s\"],\"pattern\":1}}", rgb1[2:]))
			_ = Wb.SetCellStyle("Sheet1", fmt.Sprintf("I%v", job.Id), fmt.Sprintf("I%v", job.Id), style1)
		}

		if len(rgb2) > 2 {
			style2, _ := Wb.NewStyle(fmt.Sprintf("{\"fill\":{\"type\":\"pattern\",\"color\":[\"#%s\"],\"pattern\":1}}", rgb2[2:]))
			_ = Wb.SetCellStyle("Sheet1", fmt.Sprintf("J%v", job.Id), fmt.Sprintf("J%v", job.Id), style2)
		}

		if len(rgb3) > 2 {
			style3, _ := Wb.NewStyle(fmt.Sprintf("{\"fill\":{\"type\":\"pattern\",\"color\":[\"#%s\"],\"pattern\":1}}", rgb3[2:]))
			_ = Wb.SetCellStyle("Sheet1", fmt.Sprintf("K%v", job.Id), fmt.Sprintf("K%v", job.Id), style3)
		}

		wg.Done()
	}
}
