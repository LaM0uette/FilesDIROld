package pkg

import (
	"FilesDIR/config"
	"FilesDIR/loger"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"path/filepath"
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

	if err := Wb.SaveAs(filepath.Join(config.DstPath, "exports", fmt.Sprintf("Export_%v.xlsx", time.Now().Format("20060102150405")))); err != nil {
		loger.Error("Erreur pendant la sauvegarde du fichier Excel:", err)
	}

}
