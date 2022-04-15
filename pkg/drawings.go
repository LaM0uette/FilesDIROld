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
func (s *Search) DrawStart() {
	if s.Silent {
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

func (s *Search) DrawEnd() {
	if s.Silent {
		return
	}

	defer time.Sleep(1 * time.Second)

	loger.Ui(author+config.Author, "\n", version+config.Version)

	fmt.Print(author+rgb.Green.Sprint(config.Author), "\n", version+rgb.Green.Sprint(config.Version))
	fmt.Print("\n\n")
}

func (s *Search) DrawParam(v ...any) {
	if s.Silent {
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

func (s *Search) DrawFilesOk(file string) {
	loger.Ok(fmt.Sprintf("N°%v || Fichier: %s",
		s.Counter.NbrFiles, file))
}

func (s *Search) DrawFilesSearched() {
	loger.Void(fmt.Sprintf("Traités: %v || Trouvés: %v || Dossiers: %v",
		s.Counter.NbrAllFiles, s.Counter.NbrFiles, s.Counter.NbrFolder))
}

func (s *Search) DrawBilanSearch() {
	if !s.Silent {
		defer time.Sleep(1 * time.Second)
	}

	s.DrawSep("   BILAN   ")

	loger.Ui("INFOS GENERALES:")
	loger.Ui("  Dossier principale: ", s.SrcPath)
	loger.Ui("  Requête utilisée: ", s.ReqUse)
	loger.Ui("  Nombre de Threads: ", s.Process.NbrThreads)
	loger.Ui("  Nombre de Goroutines: ", s.Process.NbrGoroutines)
	loger.Ui("\n")
	loger.Ui("RESULTATS DE LA RECHERCHE:")
	loger.Ui("  Fichiers traités: ", s.Counter.NbrAllFiles)
	loger.Ui("  Fichiers trouvés: ", s.Counter.NbrFiles)
	loger.Ui("  Temps d'exécution de la recherche: ", s.Timer.SearchEnd)
	loger.Ui("  Temps d'exécution total: ", s.Timer.AppEnd)
	loger.Ui("\n")
	loger.Ui("EXPORTS:")
	loger.Ui("  Logs: ", filepath.Join(s.DstPath, "logs"))
	loger.Ui("  Dumps: ", filepath.Join(s.DstPath, "dumps"))
	loger.Ui("  Export Excel: ", filepath.Join(s.DstPath, "exports"))

	fmt.Println(rgb.MajentaBg.Sprint("INFOS GENERALES:"))
	fmt.Println(rgb.Majenta.Sprint("  Dossier principale:"), rgb.GreenB.Sprint(s.SrcPath))
	fmt.Println(rgb.Majenta.Sprint("  Requête utilisée:"), rgb.GreenB.Sprint(s.ReqUse))
	fmt.Println(rgb.Majenta.Sprint("  Nombre de Threads:"), rgb.GreenB.Sprint(s.Process.NbrThreads))
	fmt.Println(rgb.Majenta.Sprint("  Nombre de Goroutines:"), rgb.GreenB.Sprint(s.Process.NbrGoroutines))
	fmt.Println()
	fmt.Println(rgb.MajentaBg.Sprint("RESULTATS DE LA RECHERCHE:"))
	fmt.Println(rgb.Majenta.Sprint("  Fichiers traités:"), rgb.GreenB.Sprint(s.Counter.NbrAllFiles))
	fmt.Println(rgb.Majenta.Sprint("  Fichiers trouvés:"), rgb.GreenB.Sprint(s.Counter.NbrFiles))
	fmt.Println(rgb.Majenta.Sprint("  Temps d'exécution de la recherche:"), rgb.GreenB.Sprint(s.Timer.SearchEnd))
	fmt.Println(rgb.Majenta.Sprint("  Temps d'exécution total:"), rgb.GreenB.Sprint(s.Timer.AppEnd))
	fmt.Println()
	fmt.Println(rgb.MajentaBg.Sprint("EXPORTS:"))
	fmt.Println(rgb.Majenta.Sprint("  Logs:"), rgb.GreenB.Sprint(filepath.Join(s.DstPath, "logs")))
	fmt.Println(rgb.Majenta.Sprint("  Dumps:"), rgb.GreenB.Sprint(filepath.Join(s.DstPath, "dumps")))
	fmt.Println(rgb.Majenta.Sprint("  Export Excel:"), rgb.GreenB.Sprint(filepath.Join(s.DstPath, "exports")))

	s.DrawSep("    FIN    ")
}

func (s *Search) DrawCls() {
	loger.Ok(fmt.Sprintf("Dossier %s vidé !", filepath.Base(config.LogsPath)))
	loger.Ok(fmt.Sprintf("Dossier %s vidé !", filepath.Base(config.DumpsPath)))
	loger.Ok(fmt.Sprintf("Dossier %s vidé !\n\n", filepath.Base(config.ExportsPath)))

	loger.Ui(fmt.Sprintf("Temps d'exécution: %v", s.Timer.AppEnd))
}

// ...
// Ui
func (s *Search) DrawSep(name string) {
	if s.Silent {
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
