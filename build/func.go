package build

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Search struct {
	Mode      string
	Word      string
	Path      string
	Extension string
}

// CurrentDir : Return the actual directory
func CurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func (s *Search) SearchFiles() error {

	id := 1
	DrawStartSearch()
	wb := excelize.NewFile()
	wb.SetCellValue("Sheet1", "A1", "id")
	wb.SetCellValue("Sheet1", "B1", "Fichier")
	wb.SetCellValue("Sheet1", "C1", "Date")
	wb.SetCellValue("Sheet1", "D1", "Lien_Fichier")
	wb.SetCellValue("Sheet1", "E1", "Lien")

	err := filepath.Walk(s.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() == false && strings.Contains(strings.ToLower(path), strings.ToLower(s.Word)) {

			fileStat, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("N°%v | Fichier: %v\n", id, fileStat.Name())
			wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", id+1), id)
			wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", id+1), fileStat.Name())
			wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", id+1), fmt.Sprintf("%v", fileStat.ModTime()))
			wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", id+1), path)
			wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", id+1), filepath.Dir(path))
			id++
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := wb.SaveAs("FilesDIR_Data.xlsx"); err != nil {
		log.Fatal(err)
	}

	DrawEndSearch(s.Path, "f", "g", id)
	return nil
}
