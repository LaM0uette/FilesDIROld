package build

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
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

type DataJson struct {
	Id       int    `json:"id"`
	File     string `json:"Fichier"`
	Date     string `json:"Date"`
	PathFile string `json:"Lien_Fichier"`
	Path     string `json:"Lien"`
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

	var CsvData []DataJson
	id := 0

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
		if info.IsDir() == false {

			fileStat, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}

			switch s.Mode {
			case "%":
				if !strings.Contains(strings.ToLower(fileStat.Name()), strings.ToLower(s.Word)) {
					return nil
				}
			case "=":
				if strings.Split(strings.ToLower(fileStat.Name()), ".")[0] != strings.ToLower(s.Word) {
					return nil
				}
			case "^":
				if !strings.HasPrefix(strings.ToLower(fileStat.Name()), strings.ToLower(s.Word)) {
					return nil
				}
			case "$":
				if !strings.HasSuffix(strings.Split(strings.ToLower(fileStat.Name()), ".")[0], strings.ToLower(s.Word)) {
					return nil
				}
			default:
				if !strings.Contains(strings.ToLower(fileStat.Name()), strings.ToLower(s.Word)) {
					return nil
				}
			}

			if s.Extension != "*" && strings.Split(strings.ToLower(fileStat.Name()), ".")[1] != s.Extension {
				return nil
			}

			id++

			fmt.Printf("N°%v | Fichier: %v\n", id, fileStat.Name())
			wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", id+1), id)
			wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", id+1), fileStat.Name())
			wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", id+1), fmt.Sprintf("%v", fileStat.ModTime()))
			wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", id+1), path)
			wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", id+1), filepath.Dir(path))

			data := DataJson{
				Id:       id,
				File:     fileStat.Name(),
				Date:     fmt.Sprintf("%v", fileStat.ModTime()),
				PathFile: path,
				Path:     filepath.Dir(path),
			}
			CsvData = append(CsvData, data)

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

	file, _ := json.MarshalIndent(CsvData, "", " ")
	_ = ioutil.WriteFile("FilesDIR_Data.json", file, 0644)

	DrawEndSearch(s.Path, "f", "g", id)
	return nil
}
