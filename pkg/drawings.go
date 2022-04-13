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

	sep = `~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`

	author  = `Auteur:  `
	version = `Version: `
)

func DrawStart() {
	defer time.Sleep(1 * time.Second)

	loger.Ui(start)
	loger.Ui("\t\t", author+config.Author, "\n", "\t\t", version+config.Version)
	loger.Ui("\n\n")
	loger.Ui(sep)

	rgb.HiGreen.Println(start)
	fmt.Print("\t\t", author+rgb.HiGreen.Sprint(config.Author), "\n", "\t\t", version+rgb.HiGreen.Sprint(config.Version))
	fmt.Print("\n\n")
	fmt.Println(sep)
}

func DrawEnd() {
	defer time.Sleep(1 * time.Second)

	//loger.Ui(start)
	//loger.Ui(author+config.Author, "\n", version+config.Version)

	fmt.Println(sep)
	fmt.Print(author+rgb.HiGreen.Sprint(config.Author), "\n", version+rgb.HiGreen.Sprint(config.Version))
	fmt.Print("\n\n")
}
