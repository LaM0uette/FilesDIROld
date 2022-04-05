package build

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Search struct {
	Mode       string
	Word       string
	Ext        string
	Maj        bool
	Save       bool
	Path       string
	SaveFolder string
}
type DataJson struct {
	Id       int    `json:"id"`
	File     string `json:"Fichier"`
	Date     string `json:"Date"`
	PathFile string `json:"Lien_Fichier"`
	Path     string `json:"Lien"`
}

func CurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func DesktopDir() string {
	GUID, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	homeDir := GUID.HomeDir
	return homeDir + "/Desktop/"
}

// SearchFiles : Function for search all files with different criteria in folder and sub folders
func (s *Search) SearchFiles() error {
	var searchWord string
	var JsonData []DataJson

	word := strToLower(s.Word, s.Maj)
	savePath := filepath.Join(s.SaveFolder, "Data")
	nbFolder, listFolders := countNbFolder(s.Path)
	nbFolderMade := 0
	id := 0

	err := createSaveFolder(savePath) // create folder for save data
	if err != nil {
		fmt.Println(err)
		return err
	}

	wb := excelize.NewFile()
	wb.SetCellValue("Sheet1", "A1", "id")
	wb.SetCellValue("Sheet1", "B1", "Fichier")
	wb.SetCellValue("Sheet1", "C1", "Date")
	wb.SetCellValue("Sheet1", "D1", "Lien_Fichier")
	wb.SetCellValue("Sheet1", "E1", "Lien")

	// loop for all files in folder
	err = filepath.Walk(s.Path, func(path string, info os.FileInfo, err error) error {
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

			data := DataJson{
				Id:       id,
				File:     fileStat.Name(),
				Date:     fmt.Sprintf("%v", fileStat.ModTime()),
				PathFile: path,
				Path:     filepath.Dir(path),
			}
			JsonData = append(JsonData, data)

			if s.Save {
				savelFile(wb, savePath, s.Word, JsonData)
			}
		} else {
			if stringInSlice(path, listFolders) {
				nbFolderMade++
				fmt.Printf(`
*******************************************
******        Dossier : %v/%v        ******
*******************************************

`, nbFolderMade, nbFolder)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	// save excel file
	savelFile(wb, savePath, s.Word, JsonData)

	DrawEndSearch(s.Path, savePath, nbFolderMade, id)

	return nil
}

func countNbFolder(path string) (int, []string) {
	count := 0
	var listFolders []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			count++
			listFolders = append(listFolders, path+"\\"+file.Name())
		}
	}
	return count, listFolders
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func strToLower(s string, b bool) string {
	if !b {
		return strings.ToLower(s)
	} else {
		return s
	}
}

func createSaveFolder(savePath string) error {
	err := os.MkdirAll(savePath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func savelFile(wb *excelize.File, savePath, word string, JsonData []DataJson) {
	if len(word) < 1 {
		word = "Data"
	}

	if err := wb.SaveAs(savePath + "/" + word + ".xlsx"); err != nil {
		fmt.Println(err)
	}

	file, _ := json.MarshalIndent(JsonData, "", " ")
	_ = ioutil.WriteFile(savePath+"/"+word+".json", file, 0644)
}
