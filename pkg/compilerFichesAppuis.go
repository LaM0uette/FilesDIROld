package pkg

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"path/filepath"
	"time"
)

var (
//wg   sync.WaitGroup
//jobs = make(chan compilData)
//Id   int
)

func (s *Search) CompileFichesAppuis() {
	s.DrawSep("PARAMETRES")
	s.DrawParam("INITIALISATION DE LA COMPILATION EN COURS")

	s.DrawParam("CREATION DE LA FICHE EXCEL")
	createWB()
	s.DrawSep("COMPILATION")

	if err := Wb.SaveAs(filepath.Join(GetCurrentDir(), fmt.Sprintf("__COMPILATION__%v.xlsx", time.Now().Format("20060102150405")))); err != nil {
		fmt.Println(err)
	}

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
