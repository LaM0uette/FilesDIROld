package pkg

import (
	"FilesDIR/loger"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	jobs = make(chan string)
)

func (s *Search) CompileFichesAppuis() {
	s.DrawSep("PARAMETRES")
	s.DrawParam("INITIALISATION DE LA COMPILATION EN COURS")

	s.DrawParam("CREATION DE LA FICHE EXCEL")

	createWB()

	s.DrawParam("CREATION DES WORKERS")
	for w := 1; w <= 400; w++ {
		go s.worker()
	}

	s.DrawSep("COMPILATION")

	files, err := ioutil.ReadDir(s.SrcPath)
	if err != nil {
		loger.Crash("Crash avec ce dossier:", s.SrcPath)
	}

	for _, file := range files {
		if !file.IsDir() && !strings.Contains(file.Name(), "__COMPILATION__") {

			excelFile := filepath.Join(s.SrcPath, file.Name())
			f, err := xlsx.OpenFile(excelFile)
			if err != nil {
				loger.Error("Error avec ce fichier:", excelFile)
				continue
			}

			sht := f.Sheets[0]
			maxRow := sht.MaxRow

			for i := 0; i < maxRow; i++ {
				row, err := sht.Row(i)
				if err != nil {
					loger.Crash("Crash:", err)
				}

				if i == 0 {
					continue
				}

				go func() {
					Mu.Lock()
					wg.Add(1)
					s.Counter.NbrFiles++

					jobs <- row.GetCell(3).String()
					Mu.Unlock()
				}()
			}
		}
	}

	wg.Wait()

	time.Sleep(1 * time.Second)

	s.DrawSep("EXPORT XLSX")

	if err := Wb.SaveAs(filepath.Join(s.SrcPath, fmt.Sprintf("__COMPILATION__%v.xlsx", time.Now().Format("20060102150405")))); err != nil {
		fmt.Println(err)
	}
	loger.Ok("Fichier Excel sauvegardé avec succes !")

	s.DrawSep("BILAN")
}

func createWB() {
	headers := map[string]string{
		"A1": "Chemin de la fiche",
		"B1": "Adresse",
		"C1": "Ville",
		"D1": "Num appui",
		"E1": "Type appui",
		"F1": "Type_n_app",
		"G1": "Nature TVX",
		"H1": "Etiquette jaune",
		"I1": "Effort avant ajout câble",
		"J1": "Effort après ajout câble",
		"K1": "Effort nouveau appui",
		"L1": "Latitude",
		"M1": "Longitude",
		"N1": "Opérateur",
		"O1": "Appui utilisable en l'état",
		"P1": "Environnement",
		"Q1": "Commentaire appui",
		"R1": "Commentaire global",
		"S1": "Proxi ENEDIS",
		"T1": "id_metier_",
		"U1": "Date",
		"V1": "PB",
	}

	Wb = excelize.NewFile()
	for header := range headers {
		_ = Wb.SetCellValue("Sheet1", header, headers[header])
	}
}

//...
//WORKER:
func (s *Search) worker() {
	for job := range jobs {

		excelFile := job
		f, err := xlsx.OpenFile(excelFile)
		if err != nil {
			loger.Nok(fmt.Sprintf("Error avec cette fiche appui: %s", job))

			_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", s.Counter.NbrFiles), job)

			wg.Done()
			time.Sleep(800 * time.Millisecond)
			continue
		}

		loger.Ok(fmt.Sprintf("Fiche %v ajoutée: %s", s.Counter.NbrFiles, job))

		sht := f.Sheets[0]

		adresse, _ := sht.Cell(4, 3)
		ville, _ := sht.Cell(3, 3)
		numAppui, _ := sht.Cell(2, 3)
		type1, _ := sht.Cell(25, 2)
		typeNApp, _ := sht.Cell(51, 12)
		natureTvx, _ := sht.Cell(52, 12)

		cellEtiquetteJaune, _ := sht.Cell(11, 20)
		etiquetteJaune := ""
		switch StrToLower(cellEtiquetteJaune.Value) {
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
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", s.Counter.NbrFiles), job)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", s.Counter.NbrFiles), adresse.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", s.Counter.NbrFiles), ville.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", s.Counter.NbrFiles), numAppui.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", s.Counter.NbrFiles), type1.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("F%v", s.Counter.NbrFiles), typeNApp.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("G%v", s.Counter.NbrFiles), natureTvx.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("H%v", s.Counter.NbrFiles), etiquetteJaune)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("I%v", s.Counter.NbrFiles), effort1.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("J%v", s.Counter.NbrFiles), effort2.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("K%v", s.Counter.NbrFiles), effort3.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("L%v", s.Counter.NbrFiles), lat.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("M%v", s.Counter.NbrFiles), lon.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("N%v", s.Counter.NbrFiles), operateur.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("O%v", s.Counter.NbrFiles), utilisableEnEtat.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("P%v", s.Counter.NbrFiles), environnement.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("Q%v", s.Counter.NbrFiles), commentaireEtatAppui.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("R%v", s.Counter.NbrFiles), commentaireGlobal.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("S%v", s.Counter.NbrFiles), proxiEnedis.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("T%v", s.Counter.NbrFiles), idMetier)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("U%v", s.Counter.NbrFiles), date.Value)
		_ = Wb.SetCellValue("Sheet1", fmt.Sprintf("V%v", s.Counter.NbrFiles), pb.Value)

		rgb1 := effort1.GetStyle().Fill.FgColor
		rgb2 := effort2.GetStyle().Fill.FgColor
		rgb3 := effort3.GetStyle().Fill.FgColor

		if len(rgb1) > 2 {
			style1, _ := Wb.NewStyle(fmt.Sprintf("{\"fill\":{\"type\":\"pattern\",\"color\":[\"#%s\"],\"pattern\":1}}", rgb1[2:]))
			_ = Wb.SetCellStyle("Sheet1", fmt.Sprintf("I%v", s.Counter.NbrFiles), fmt.Sprintf("I%v", s.Counter.NbrFiles), style1)
		}

		if len(rgb2) > 2 {
			style2, _ := Wb.NewStyle(fmt.Sprintf("{\"fill\":{\"type\":\"pattern\",\"color\":[\"#%s\"],\"pattern\":1}}", rgb2[2:]))
			_ = Wb.SetCellStyle("Sheet1", fmt.Sprintf("J%v", s.Counter.NbrFiles), fmt.Sprintf("J%v", s.Counter.NbrFiles), style2)
		}

		if len(rgb3) > 2 {
			style3, _ := Wb.NewStyle(fmt.Sprintf("{\"fill\":{\"type\":\"pattern\",\"color\":[\"#%s\"],\"pattern\":1}}", rgb3[2:]))
			_ = Wb.SetCellStyle("Sheet1", fmt.Sprintf("K%v", s.Counter.NbrFiles), fmt.Sprintf("K%v", s.Counter.NbrFiles), style3)
		}

		wg.Done()
	}
}
