package pkg

import (
	"FilesDIR/config"
	"FilesDIR/loger"
	"FilesDIR/rgb"
	"fmt"
	"path/filepath"
	"time"
)

const (
	start = `
		███████╗██╗██╗     ███████╗██████╗ ██╗██████╗ 
		██╔════╝██║██║     ██╔════╝██╔══██╗██║██╔══██╗
		█████╗  ██║██║     █████╗  ██║  ██║██║██████╔╝
		██╔══╝  ██║██║     ██╔══╝  ██║  ██║██║██╔══██╗
		██║     ██║███████╗███████╗██████╔╝██║██║  ██║
		╚═╝     ╚═╝╚══════╝╚══════╝╚═════╝ ╚═╝╚═╝  ╚═╝`
	ligneSep = `■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■`

	author  = `Auteur:  `
	version = `Version: `
)

// ...
// FilesDIR
func (flg *SSearch) DrawStart() {
	if flg.Silent {
		return
	}

	defer time.Sleep(1 * time.Second)

	loger.Ui(start)
	loger.Ui("\t\t", author+config.Author, "\n", "\t\t", version+config.Version)
	loger.Ui("\n")

	rgb.Green.Println(start)
	fmt.Print("\t\t", author+rgb.Green.Sprint(config.Author), "\n", "\t\t", version+rgb.Green.Sprint(config.Version))
	fmt.Print("\n\n")
}

func (flg *SSearch) DrawEnd() {
	if flg.Silent {
		return
	}

	defer time.Sleep(1 * time.Second)

	loger.Ui(author+config.Author, "\n", version+config.Version)

	fmt.Print(author+rgb.Green.Sprint(config.Author), "\n", version+rgb.Green.Sprint(config.Version))
	fmt.Print("\n\n")
}

func (flg *SSearch) DrawParam(v ...any) {
	if flg.Silent {
		return
	}

	defer time.Sleep(400 * time.Millisecond)

	pre := "██████████"
	txt := ""
	arg1 := ""
	arg2 := ""

	if len(v) >= 1 {
		txt = fmt.Sprintf(" %s", v[0])
	}
	if len(v) >= 2 {
		arg1 = fmt.Sprintf(" %s", v[1])
	}
	if len(v) >= 3 {
		arg2 = fmt.Sprintf(" %s", v[2])
	}

	loger.Ui(pre, txt, arg1, arg2)

	fmt.Printf("%s%s%s%s\n", rgb.YellowUB.Sprint(pre),
		rgb.YellowUB.Sprint(txt), rgb.GreenB.Sprint(arg1), rgb.Gray.Sprint(arg2))
}

func (flg *SSearch) DrawFilesOk(file string) {

	if len(file) > 40 {
		loger.Ok(fmt.Sprintf("N°%v || Fichier: %s",
			flg.Counter.NbrFiles, file))
	} else {
		loger.Ok(fmt.Sprintf("N°%v || Fichier: %s                                           ",
			flg.Counter.NbrFiles, file))
	}
}

func (flg *SSearch) DrawFilesSearched() {
	loger.Void(fmt.Sprintf("Traités: %v || Trouvés: %v || Dossiers: %v",
		flg.Counter.NbrAllFiles, flg.Counter.NbrFiles, flg.Counter.NbrFolder))
}

func (flg *SSearch) DrawBilanSearch() {
	if !flg.Silent {
		defer time.Sleep(1 * time.Second)
	}

	flg.DrawSep("   BILAN   ")

	loger.Ui("INFOS GENERALES:")
	loger.Ui("  Dossier principale: ", flg.SrcPath)
	loger.Ui("  Requête utilisée: ", flg.ReqUse)
	loger.Ui("  Nombre de Threads: ", flg.Process.NbrThreads)
	loger.Ui("  Nombre de Goroutines: ", flg.Process.NbrGoroutines)
	loger.Ui("\n")
	loger.Ui("RESULTATS DE LA RECHERCHE:")
	loger.Ui("  Fichiers traités: ", flg.Counter.NbrAllFiles)
	loger.Ui("  Fichiers trouvés: ", flg.Counter.NbrFiles)
	loger.Ui("  Temps d'exécution de la recherche: ", flg.Timer.SearchEnd)
	loger.Ui("  Temps d'exécution total: ", flg.Timer.AppEnd)
	loger.Ui("\n")
	loger.Ui("EXPORTS:")
	loger.Ui("  Logs: ", filepath.Join(flg.DstPath, "logs"))
	loger.Ui("  Dumps: ", filepath.Join(flg.DstPath, "dumps"))
	loger.Ui("  Export Excel: ", filepath.Join(flg.DstPath, "exports"))

	fmt.Println(rgb.MajentaBg.Sprint("INFOS GENERALES:"))
	fmt.Println(rgb.Majenta.Sprint("  Dossier principale:"), rgb.GreenB.Sprint(flg.SrcPath))
	fmt.Println(rgb.Majenta.Sprint("  Requête utilisée:"), rgb.GreenB.Sprint(flg.ReqUse))
	fmt.Println(rgb.Majenta.Sprint("  Nombre de Threads:"), rgb.GreenB.Sprint(flg.Process.NbrThreads))
	fmt.Println(rgb.Majenta.Sprint("  Nombre de Goroutines:"), rgb.GreenB.Sprint(flg.Process.NbrGoroutines))
	fmt.Println()
	fmt.Println(rgb.MajentaBg.Sprint("RESULTATS DE LA RECHERCHE:"))
	fmt.Println(rgb.Majenta.Sprint("  Fichiers traités:"), rgb.GreenB.Sprint(flg.Counter.NbrAllFiles))
	fmt.Println(rgb.Majenta.Sprint("  Fichiers trouvés:"), rgb.GreenB.Sprint(flg.Counter.NbrFiles))
	fmt.Println(rgb.Majenta.Sprint("  Temps d'exécution de la recherche:"), rgb.GreenB.Sprint(flg.Timer.SearchEnd))
	fmt.Println(rgb.Majenta.Sprint("  Temps d'exécution total:"), rgb.GreenB.Sprint(flg.Timer.AppEnd))
	fmt.Println()
	fmt.Println(rgb.MajentaBg.Sprint("EXPORTS:"))
	fmt.Println(rgb.Majenta.Sprint("  Logs:"), rgb.GreenB.Sprint(filepath.Join(flg.DstPath, "logs")))
	fmt.Println(rgb.Majenta.Sprint("  Dumps:"), rgb.GreenB.Sprint(filepath.Join(flg.DstPath, "dumps")))
	fmt.Println(rgb.Majenta.Sprint("  Export Excel:"), rgb.GreenB.Sprint(filepath.Join(flg.DstPath, "exports")))

	flg.DrawSep("    FIN    ")
}

func (flg *SSearch) DrawCls() {
	loger.Ok(fmt.Sprintf("Dossier %s vidé !", filepath.Base(config.LogsPath)))
	loger.Ok(fmt.Sprintf("Dossier %s vidé !", filepath.Base(config.DumpsPath)))
	loger.Ok(fmt.Sprintf("Dossier %s vidé !\n", filepath.Base(config.ExportsPath)))

	loger.Ui(fmt.Sprintf("\nTemps d'exécution: %v\n", flg.Timer.AppEnd))
	loger.Void(fmt.Sprintf("Temps d'exécution: %v\n\n", flg.Timer.AppEnd))
}

func (flg *SSearch) DrawFichesAppuisCompiled() {
	loger.Ui(fmt.Sprintf("%v fiches compilées avec succes !", flg.Counter.NbrFiles))
	loger.Void(fmt.Sprintf("%v fiches compilées avec succes !\n", flg.Counter.NbrFiles))

	loger.Ui(fmt.Sprintf("Temps d'exécution: %v\n\n", flg.Timer.AppEnd))
	loger.Void(fmt.Sprintf("Temps d'exécution: %v\n\n", flg.Timer.AppEnd))
}

// ...
// Ui
func (flg *SSearch) DrawSep(name string) {
	if flg.Silent {
		return
	}

	sep := ligneSep + fmt.Sprintf(" %s ", name) + ligneSep
	sepRgb := rgb.Gray.Sprint(ligneSep) + rgb.GreenB.Sprintf(" %s ", name) + rgb.Gray.Sprint(ligneSep)

	loger.Ui("\n\n", sep)
	fmt.Println("\n\n", sepRgb)
}

// ...
// Export Excel
func DrawAddLine(job, iMax int) {
	loger.Void(fmt.Sprintf("Export Excel: %v/%v", job, iMax))
}
