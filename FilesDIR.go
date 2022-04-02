package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Data struct {
	Id   int    `json:"id"`
	Name string `json:"nom"`
	Date string `json:"date"`
	Path string `json:"lien"`
}

var reader = bufio.NewReader(os.Stdin)

func convertJSONToCSV(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	var ranking []Data
	if err := json.NewDecoder(sourceFile).Decode(&ranking); err != nil {
		return err
	}

	// 3. Create a new file to store CSV data
	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 4. Write the header of the CSV file and the successive rows by iterating through the JSON struct array
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"Id", "Name", "Date", "Path"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range ranking {
		var csvRow []string
		csvRow = append(csvRow, fmt.Sprint(r.Id), r.Name, fmt.Sprint(r.Date), r.Path)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "B2", 100)
	f.SetCellValue("Sheet1", "A1", 50)

	now := time.Now()

	f.SetCellValue("Sheet1", "A4", now.Format(time.ANSIC))

	if err := f.SaveAs("simple.xlsx"); err != nil {
		log.Fatal(err)
	}

	var CsvData []Data
	var id = 0

	fmt.Print("\nMot clé : ")
	mot, err := reader.ReadString('\n')
	mot = strings.TrimSpace(mot)

	fmt.Print("\nLien de recherche : ")
	lienRe, err := reader.ReadString('\n')
	lienRe = strings.TrimSpace(lienRe)

	fmt.Println("**************  Début de la recherche  **************")
	err = filepath.Walk(lienRe, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() == false && strings.Contains(strings.ToLower(path), strings.ToLower(mot)) {

			fileStat, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Nom: %v | Id: %v\n", fileStat.Name(), id)

			data := Data{
				Id:   id,
				Name: fileStat.Name(),
				Date: fmt.Sprintf("%v", fileStat.ModTime()),
				Path: path,
			}
			CsvData = append(CsvData, data)
			id++

		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("**************  Recherche terminée  **************\n")

	fmt.Println("\n**************  Création du json  **************")
	file, _ := json.MarshalIndent(CsvData, "", " ")
	_ = ioutil.WriteFile("data.json", file, 0644)
	fmt.Println("**************  Création du json terminé  **************\n")

	fmt.Println("\n**************  Création du csv  **************")
	if err := convertJSONToCSV("data.json", "data.csv"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("**************  Création du csv terminé  **************\n")
}
