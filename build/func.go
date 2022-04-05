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
	"strconv"
	"strings"
	"time"
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
func (s *Search) SearchFiles() (reqUse, savePath string, nbFolderMade, id int, err error) {
	var searchWord string
	var JsonData []DataJson

	word := strToLower(s.Word, s.Maj)
	savePath = filepath.Join(s.SaveFolder, "Data")
	nbFolder, listFolders := countFoldersDir(s.Path)
	nbFolderMade = 0
	id = 0

	sMaj := ""
	if s.Maj {
		sMaj = " -maj"
	}
	sSave := ""
	if s.Save {
		sSave = " -save"
	}
	reqUse = fmt.Sprintf("FilesDIR -r -mode=%s -word=%s -ext=%s%s%s\n", s.Mode, s.Word, s.Ext, sMaj, sSave)

	err = createSaveFolder(savePath) // create folder for save data
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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

			go fmt.Printf("N°%v | Fichier: %v\n", id, fileStat.Name())

			wb.SetCellValue("Sheet1", fmt.Sprintf("A%v", id+1), id)
			wb.SetCellValue("Sheet1", fmt.Sprintf("B%v", id+1), fileStat.Name())
			wb.SetCellValue("Sheet1", fmt.Sprintf("C%v", id+1), fmt.Sprintf("%v", fileStat.ModTime()))
			wb.SetCellValue("Sheet1", fmt.Sprintf("D%v", id+1), path)
			wb.SetCellValue("Sheet1", fmt.Sprintf("E%v", id+1), filepath.Dir(path))

			JsonData = addJsonData(path, id, fileStat, JsonData)

			if s.Save {
				savelFiles(wb, savePath, s.Word, JsonData, false)
			}
		} else {
			if stringInSlice(path, listFolders) {
				savelFiles(wb, savePath, s.Word, JsonData, false)
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
	}

	// save excel file
	savelFiles(wb, savePath, s.Word, JsonData, true)

	return reqUse, savePath, nbFolderMade, id, nil
}

func addJsonData(path string, id int, fileStat os.FileInfo, JsonData []DataJson) []DataJson {
	data := DataJson{
		Id:       id,
		File:     fileStat.Name(),
		Date:     fmt.Sprintf("%v", fileStat.ModTime()),
		PathFile: path,
		Path:     filepath.Dir(path),
	}
	JsonData = append(JsonData, data)
	return JsonData
}

func countFoldersDir(path string) (int, []string) {
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

func strToLower(s string, b bool) string {
	if !b {
		return strings.ToLower(s)
	} else {
		return s
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func createSaveFolder(savePath string) error {
	err := os.MkdirAll(savePath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func savelFiles(wb *excelize.File, savePath, word string, JsonData []DataJson, end bool) {
	if len(word) < 1 {
		word = "Data"
	}

	if err := wb.SaveAs(savePath + "\\" + word + ".xlsx"); err != nil {
		fmt.Println(err)
	}
	if end {
		currentTime := time.Now()

		if err := wb.SaveAs("T:\\- 4 Suivi Appuis\\26_MACROS\\GO\\FilesDIR\\Data\\" + word + "__" + currentTime.Format("2017-09-07 2:3:5") + ".xlsx"); err != nil {
			fmt.Println(err)
		}
	}

	file, _ := json.MarshalIndent(JsonData, "", " ")
	_ = ioutil.WriteFile(savePath+"/"+word+".json", file, 0644)
}

func ReadExcelFileForReq(req string) (mode, word, ext string, maj, save bool, err error) {
	f, err := excelize.OpenFile("T:\\- 4 Suivi Appuis\\26_MACROS\\GO\\FilesDIR\\req.xlsx")
	if err != nil {
		panic(err)
	}

	i := 1
	for len(f.GetCellValue("Sheet1", fmt.Sprintf("A%v", i))) > 0 {
		if strings.ToLower(req) == strings.ToLower(f.GetCellValue("Sheet1", fmt.Sprintf("A%v", i))) {
			break
		}
		i++
	}
	if len(f.GetCellValue("Sheet1", fmt.Sprintf("A%v", i))) < 1 {
		os.Exit(1)
	}

	mode = f.GetCellValue("Sheet1", fmt.Sprintf("B%v", i))
	word = f.GetCellValue("Sheet1", fmt.Sprintf("C%v", i))
	ext = f.GetCellValue("Sheet1", fmt.Sprintf("D%v", i))
	strMaj := f.GetCellValue("Sheet1", fmt.Sprintf("E%v", i))
	strSave := f.GetCellValue("Sheet1", fmt.Sprintf("F%v", i))

	maj, err = strconv.ParseBool(strMaj)
	save, err = strconv.ParseBool(strSave)

	return mode, word, ext, maj, save, nil
}
