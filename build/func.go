package build

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// Search : Structure with data to search for files
type Search struct {
	Mode       string
	Word       string
	Ext        string
	Maj        bool
	Path       string
	SaveFolder string
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

// DesktopDir : Return the desktop directory
func DesktopDir() string {
	myself, err := user.Current()
	if err != nil {
		panic(err)
	}
	homedir := myself.HomeDir
	desktop := homedir + "/Desktop/"
	return desktop
}

// SearchFiles : Function for search all files with different criteria in folder and sub folders
func (s *Search) SearchFiles() error {
	word := strToLower(s.Word, s.Maj)
	var searchWord string

	var JsonData []DataJson // Var for generate json file
	id := 0                 // count number file searched

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
				return err
			}

			searchWord = strToLower(fileStat.Name(), s.Maj)

			// condition of search Mode ( = | % | ^ | $ )
			switch s.Mode {
			case "%":
				if !strings.Contains(searchWord, word) {
					return nil
				}
			case "=":
				if strings.Split(searchWord, ".")[0] != word {
					return nil
				}
			case "^":
				if !strings.HasPrefix(searchWord, word) {
					return nil
				}
			case "$":
				if !strings.HasSuffix(strings.Split(searchWord, ".")[0], word) {
					return nil
				}
			default:
				if !strings.Contains(searchWord, word) {
					return nil
				}
			}

			// condition of extension file
			if s.Ext != "*" && strings.Split(searchWord, ".")[1] != s.Ext {
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
			JsonData = append(JsonData, data)

		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create a new folder to save data file
	savePath := s.SaveFolder + "/Data"
	err = os.MkdirAll(savePath, os.ModePerm)
	if err != nil {
		return err
	}

	// save excel file
	if err := wb.SaveAs(savePath + "/" + s.Word + ".xlsx"); err != nil {
		return err
	}

	// Generate a json file
	file, _ := json.MarshalIndent(JsonData, "", " ")
	_ = ioutil.WriteFile(savePath+"/"+s.Word+".json", file, 0644)

	// print on the screen the end of search
	DrawEndSearch(s.Path, savePath, id)
	return nil
}

// strToLower : Function for convert string to lower
func strToLower(s string, b bool) string {
	if !b {
		return strings.ToLower(s)
	} else {
		return s
	}
}
