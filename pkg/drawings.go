package pkg

import (
	"FilesDIR/config"
	"FilesDIR/loger"
	"FilesDIR/rgb"
	"fmt"
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
func DrawStart() {
	defer time.Sleep(1 * time.Second)

	loger.Ui(start)
	loger.Ui("\t\t", author+config.Author, "\n", "\t\t", version+config.Version)
	loger.Ui("\n")

	rgb.Green.Println(start)
	fmt.Print("\t\t", author+rgb.Green.Sprint(config.Author), "\n", "\t\t", version+rgb.Green.Sprint(config.Version))
	fmt.Print("\n\n")
}

func DrawEnd() {
	defer time.Sleep(1 * time.Second)

	loger.Ui(author+config.Author, "\n", version+config.Version)

	fmt.Print(author+rgb.Green.Sprint(config.Author), "\n", version+rgb.Green.Sprint(config.Version))
	fmt.Print("\n\n")
}

func DrawParam(v ...any) {
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

// ...
// Ui
func DrawSep(name string) {
	sep := ligneSep + fmt.Sprintf(" %s ", name) + ligneSep
	sepRgb := rgb.Gray.Sprint(ligneSep) + rgb.GreenB.Sprintf(" %s ", name) + rgb.Gray.Sprint(ligneSep)

	loger.Ui("\n\n", sep)
	fmt.Println("\n\n", sepRgb)
}
