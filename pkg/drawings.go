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
	sep = `~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`

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
	loger.Ui(sep)
	loger.Ui("\n")

	rgb.Green.Println(start)
	fmt.Print("\t\t", author+rgb.Green.Sprint(config.Author), "\n", "\t\t", version+rgb.Green.Sprint(config.Version))
	fmt.Print("\n\n")
	rgb.Gray.Println(sep)
	fmt.Print("\n")
}

func DrawEnd() {
	defer time.Sleep(1 * time.Second)

	loger.Ui(sep)
	loger.Ui(author+config.Author, "\n", version+config.Version)

	rgb.Gray.Println(sep)
	fmt.Print(author+rgb.Green.Sprint(config.Author), "\n", version+rgb.Green.Sprint(config.Version))
	fmt.Print("\n\n")
}

// ...
// Search
func DrawParam(v ...any) {
	defer time.Sleep(600 * time.Millisecond)

	pre := "          "
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

	fmt.Printf("%s%s%s%s\n", rgb.BgYellow.Sprint(pre),
		rgb.Yellow.Sprint(txt), rgb.GreenB.Sprint(arg1), rgb.Gray.Sprint(arg2))
}
