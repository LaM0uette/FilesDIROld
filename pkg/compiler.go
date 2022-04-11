package pkg

import (
	"FilesDIR/display"
	"FilesDIR/globals"
	"FilesDIR/loger"
	"FilesDIR/task"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type compilData struct {
	Path string
	Id   int
}

type DataExp struct {
	path                 string
	id                   int
	adresse              string
	ville                string
	numAppui             string
	type1                string
	typeNApp             string
	natureTvx            string
	etiquetteJaune       string
	effort1              string
	effort2              string
	effort3              string
	lat                  string
	lon                  string
	operateur            string
	utilisableEnEtat     string
	environnement        string
	commentaireEtatAppui string
	commentaireGlobal    string
	proxiEnedis          string
	idMetier             string
	date                 string
	pb                   string
}

var (
	wg         sync.WaitGroup
	wgWritter  sync.WaitGroup
	jobs       = make(chan compilData)
	jobsWrtier = make(chan int)
	Wb         = &xlsx.File{}
	Sht        = &xlsx.Sheet{}
	Data       []DataExp
	Id         int
)

//...
// ACTIONS:
func ClsTempFiles() {
	_ = os.RemoveAll(globals.FolderLogs)
	_ = os.RemoveAll(globals.FolderDumps)
	_ = os.RemoveAll(globals.FolderExports)
}

func CompilerFicheAppuiFt(path string) {
	path = "C:\\Users\\doria\\FilesDIR\\Nouveau dossier"

	loger.BlankDateln(display.DrawInitCompiler())
	time.Sleep(800 * time.Millisecond)

	loger.Blankln(display.DrawRunCompiler())

	Id = 1

	Wb = xlsx.NewFile()
	Sht, _ = Wb.AddSheet("Sheet1")

	col0, _ := Sht.Cell(0, 0)
	col0.SetValue("Chemin de la fiche")
	col1, _ := Sht.Cell(0, 1)
	col1.SetValue("Adresse")
	col2, _ := Sht.Cell(0, 2)
	col2.SetValue("Ville")
	col3, _ := Sht.Cell(0, 3)
	col3.SetValue("Num appui")
	col4, _ := Sht.Cell(0, 4)
	col4.SetValue("Type appui")
	col5, _ := Sht.Cell(0, 5)
	col5.SetValue("Type_n_app")
	col6, _ := Sht.Cell(0, 6)
	col6.SetValue("Nature TVX")
	col7, _ := Sht.Cell(0, 7)
	col7.SetValue("Etiquette jaune")
	col8, _ := Sht.Cell(0, 8)
	col8.SetValue("Effort avant ajout câble")
	col9, _ := Sht.Cell(0, 9)
	col9.SetValue("Effort après ajout câble")
	col10, _ := Sht.Cell(0, 10)
	col10.SetValue("Effort nouveau appui")
	col11, _ := Sht.Cell(0, 11)
	col11.SetValue("Latitude")
	col12, _ := Sht.Cell(0, 12)
	col12.SetValue("Longitude")
	col13, _ := Sht.Cell(0, 13)
	col13.SetValue("Opérateur")
	col14, _ := Sht.Cell(0, 14)
	col14.SetValue("Appui utilisable en l'état")
	col15, _ := Sht.Cell(0, 15)
	col15.SetValue("Environnement")
	col16, _ := Sht.Cell(0, 16)
	col16.SetValue("Commentaire appui")
	col17, _ := Sht.Cell(0, 17)
	col17.SetValue("Commentaire global")
	col18, _ := Sht.Cell(0, 18)
	col18.SetValue("Proxi ENEDIS")
	col19, _ := Sht.Cell(0, 19)
	col19.SetValue("id_metier_")
	col20, _ := Sht.Cell(0, 20)
	col20.SetValue("Date")
	col21, _ := Sht.Cell(0, 21)
	col21.SetValue("PB")

	for w := 1; w <= 10; w++ {
		go workerFicheAppuiFt()
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		loger.Crashln(fmt.Sprintf("Crash with this path: %s", path))
	}

	for _, file := range files {
		if !file.IsDir() && !strings.Contains(file.Name(), "__COMPILATION__") {

			excelFile := filepath.Join(path, file.Name())
			f, err := xlsx.OpenFile(excelFile)
			if err != nil {
				loger.Errorln(fmt.Sprintf("Crash with this files: %s", excelFile))
				continue
			}

			sht := f.Sheets[0]

			maxRow := sht.MaxRow

			for i := 0; i < maxRow; i++ {
				row, err := sht.Row(i)
				if err != nil {
					panic(err)
				}

				go func() {
					wg.Add(1)
					Id++

					jobs <- compilData{
						Path: row.GetCell(3).String(),
						Id:   Id,
					}
				}()
			}
		}
	}

	wg.Wait()
	time.Sleep(1 * time.Second)

	iMax := len(Data)
	for w := 1; w <= 10; w++ {
		go writeExcelLineWorker()
	}
	// Run writing loop
	for i := 0; i < iMax; i++ {
		i := i
		go func() {
			wgWritter.Add(1)
			jobsWrtier <- i
		}()
	}

	wgWritter.Wait()

	loger.BlankDateln(display.DrawEndCompiler())

	loger.BlankDateln(fmt.Sprintf("Nombre de fiches compilées : %v", Id-1))
	time.Sleep(800 * time.Millisecond)

	if err := Wb.Save(filepath.Join(path, fmt.Sprintf("__COMPILATION__%v.xlsx", time.Now().Format("20060102150405")))); err != nil {
		fmt.Println(err)
	}

	loger.Blankln(display.DrawSaveExcel())
	fmt.Println()
	time.Sleep(200 * time.Millisecond)
}

//...
//WORKER:
func workerFicheAppuiFt() {
	for job := range jobs {
		loger.BlankDateln(fmt.Sprintf("N°%v | Files: %s", job.Id, filepath.Base(job.Path)))

		excelFile := job.Path
		f, err := xlsx.OpenFile(excelFile)
		if err != nil {
			loger.Errorln(fmt.Sprintf("Crash with this files: %s", filepath.Base(excelFile)))
			wg.Done()
			continue
		}

		sht := f.Sheets[0]

		adresse, _ := sht.Cell(4, 3)
		ville, _ := sht.Cell(3, 3)
		numAppui, _ := sht.Cell(2, 3)
		type1, _ := sht.Cell(25, 2)
		typeNApp, _ := sht.Cell(51, 12)
		natureTvx, _ := sht.Cell(52, 12)

		cellEtiquetteJaune, _ := sht.Cell(11, 20)
		etiquetteJaune := ""
		switch task.StrToLower(cellEtiquetteJaune.Value) {
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

		Data = append(Data, DataExp{
			path:                 job.Path,
			id:                   job.Id,
			adresse:              adresse.Value,
			ville:                ville.Value,
			numAppui:             numAppui.Value,
			type1:                type1.Value,
			typeNApp:             typeNApp.Value,
			natureTvx:            natureTvx.Value,
			etiquetteJaune:       etiquetteJaune,
			effort1:              effort1.Value,
			effort2:              effort2.Value,
			effort3:              effort3.Value,
			lat:                  lat.Value,
			lon:                  lon.Value,
			operateur:            operateur.Value,
			utilisableEnEtat:     utilisableEnEtat.Value,
			environnement:        environnement.Value,
			commentaireEtatAppui: commentaireEtatAppui.Value,
			commentaireGlobal:    commentaireGlobal.Value,
			proxiEnedis:          proxiEnedis.Value,
			idMetier:             idMetier,
			date:                 date.Value,
			pb:                   pb.Value,
		})

		// insert value
		//col0, _ := Sht.Cell(job.Id, 0)
		//col0.SetValue(job.Path)
		//col1, _ := Sht.Cell(job.Id, 1)
		//col1.SetValue(adresse.Value)
		//col2, _ := Sht.Cell(job.Id, 2)
		//col2.SetValue(ville.Value)
		//col3, _ := Sht.Cell(job.Id, 3)
		//col3.SetValue(numAppui.Value)
		//col4, _ := Sht.Cell(job.Id, 4)
		//col4.SetValue(type1.Value)
		//col5, _ := Sht.Cell(job.Id, 5)
		//col5.SetValue(typeNApp.Value)
		//col6, _ := Sht.Cell(job.Id, 6)
		//col6.SetValue(natureTvx.Value)
		//col7, _ := Sht.Cell(job.Id, 7)
		//col7.SetValue(etiquetteJaune)
		//col8, _ := Sht.Cell(job.Id, 8)
		//col8.SetValue(effort1.Value)
		//col9, _ := Sht.Cell(job.Id, 9)
		//col9.SetValue(effort2.Value)
		//col10, _ := Sht.Cell(job.Id, 10)
		//col10.SetValue(effort3.Value)
		//col11, _ := Sht.Cell(job.Id, 11)
		//col11.SetValue(lat.Value)
		//col12, _ := Sht.Cell(job.Id, 12)
		//col12.SetValue(lon.Value)
		//col13, _ := Sht.Cell(job.Id, 13)
		//col13.SetValue(operateur.Value)
		//col14, _ := Sht.Cell(job.Id, 14)
		//col14.SetValue(utilisableEnEtat.Value)
		//col15, _ := Sht.Cell(job.Id, 15)
		//col15.SetValue(environnement.Value)
		//col16, _ := Sht.Cell(job.Id, 16)
		//col16.SetValue(commentaireEtatAppui.Value)
		//col17, _ := Sht.Cell(job.Id, 17)
		//col17.SetValue(commentaireGlobal.Value)
		//col18, _ := Sht.Cell(job.Id, 18)
		//col18.SetValue(proxiEnedis.Value)
		//col19, _ := Sht.Cell(job.Id, 19)
		//col19.SetValue(idMetier)
		//col20, _ := Sht.Cell(job.Id, 20)
		//col20.SetValue(date.Value)
		//col21, _ := Sht.Cell(job.Id, 21)
		//col21.SetValue(pb.Value)

		wg.Done()
	}
}

func writeExcelLineWorker() {
	for jobb := range jobsWrtier {

		fmt.Print("\r")
		fmt.Printf("\rSauvegarde du fichier Excel...  %v/%v", jobb, len(Data))

		col0, _ := Sht.Cell(jobb, 0)
		col0.SetValue(Data[jobb].path)
		col1, _ := Sht.Cell(jobb, 1)
		col1.SetValue(Data[jobb].adresse)
		col2, _ := Sht.Cell(jobb, 2)
		col2.SetValue(Data[jobb].ville)
		//col3, _ := Sht.Cell(job, 3)
		//col3.SetValue(numAppui.Value)
		//col4, _ := Sht.Cell(job, 4)
		//col4.SetValue(type1.Value)
		//col5, _ := Sht.Cell(job, 5)
		//col5.SetValue(typeNApp.Value)
		//col6, _ := Sht.Cell(job, 6)
		//col6.SetValue(natureTvx.Value)
		//col7, _ := Sht.Cell(job, 7)
		//col7.SetValue(etiquetteJaune)
		//col8, _ := Sht.Cell(job, 8)
		//col8.SetValue(effort1.Value)
		//col9, _ := Sht.Cell(job, 9)
		//col9.SetValue(effort2.Value)
		//col10, _ := Sht.Cell(job, 10)
		//col10.SetValue(effort3.Value)
		//col11, _ := Sht.Cell(job, 11)
		//col11.SetValue(lat.Value)
		//col12, _ := Sht.Cell(job, 12)
		//col12.SetValue(lon.Value)
		//col13, _ := Sht.Cell(job, 13)
		//col13.SetValue(operateur.Value)
		//col14, _ := Sht.Cell(job, 14)
		//col14.SetValue(utilisableEnEtat.Va
		//col15, _ := Sht.Cell(job, 15)
		//col15.SetValue(environnement.Value
		//col16, _ := Sht.Cell(job, 16)
		//col16.SetValue(commentaireEtatAppu
		//col17, _ := Sht.Cell(job, 17)
		//col17.SetValue(commentaireGlobal.V
		//col18, _ := Sht.Cell(job, 18)
		//col18.SetValue(proxiEnedis.Value)
		//col19, _ := Sht.Cell(job, 19)
		//col19.SetValue(idMetier)
		//col20, _ := Sht.Cell(job, 20)
		//col20.SetValue(date.Value)
		//col21, _ := Sht.Cell(job, 21)
		//col21.SetValue(pb.Value)

		wgWritter.Done()
	}
}
