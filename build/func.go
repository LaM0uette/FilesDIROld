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

// Search : Structure with data to search for files
type Search struct {
	Mode      string
	Word      string
	Path      string
	Extension string
}

// DataJson : Struct for generate json file
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

// SearchFiles : Function for search all files with different criteria in folder and sub folders
func (s *Search) SearchFiles() error {

	var CsvData []DataJson // Var for generate json file
	id := 0                // count number file searched

	DrawStartSearch() // print on the screen the start of search

	// Create new excel file
	wb := excelize.NewFile()
	wb.SetCellValue("Sheet1", "A1", "id")           // insert column name in excel file (A1)
	wb.SetCellValue("Sheet1", "B1", "Fichier")      // insert column name in excel file (B1)
	wb.SetCellValue("Sheet1", "C1", "Date")         // insert column name in excel file (C1)
	wb.SetCellValue("Sheet1", "D1", "Lien_Fichier") // insert column name in excel file (D1)
	wb.SetCellValue("Sheet1", "E1", "Lien")         // insert column name in excel file (E1)

	// loop for all files in folder
	err := filepath.Walk(s.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		// check if is not a folder
		if info.IsDir() == false {

			// look at the stats of file
			fileStat, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}

			// condition of search Mode ( = | % | ^ | $ )
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

			// condition of extension file
			if s.Extension != "*" && strings.Split(strings.ToLower(fileStat.Name()), ".")[1] != s.Extension {
				return nil
			}

			id++ // increment the id (number of file searched)

			fmt.Printf("N°%v | Fichier: %v\n", id, fileStat.Name()) // print searched file on screen

			wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", id+1), id)                                    // insert data of excel file
			wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", id+1), fileStat.Name())                       // insert data of excel file
			wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", id+1), fmt.Sprintf("%v", fileStat.ModTime())) // insert data of excel file
			wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", id+1), path)                                  // insert data of excel file
			wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", id+1), filepath.Dir(path))                    // insert data of excel file

			// Add entry of DataJson struct
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

	// save excel file
	if err := wb.SaveAs("FilesDIR_Data.xlsx"); err != nil {
		log.Fatal(err)
	}

	// Generate a json file
	file, _ := json.MarshalIndent(CsvData, "", " ")
	_ = ioutil.WriteFile("FilesDIR_Data.json", file, 0644)

	// print on the screen the end of search
	DrawEndSearch(s.Path, "f", "g", id)
	return nil
}
