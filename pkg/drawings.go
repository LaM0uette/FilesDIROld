package pkg

import (
	"FilesDIR/config"
	"FilesDIR/loger"
	"FilesDIR/rgb"
	"time"
)

const (
	start = `
		███████╗██╗██╗     ███████╗██████╗ ██╗██████╗ 
		██╔════╝██║██║     ██╔════╝██╔══██╗██║██╔══██╗
		█████╗  ██║██║     █████╗  ██║  ██║██║██████╔╝
		██╔══╝  ██║██║     ██╔══╝  ██║  ██║██║██╔══██╗
		██║     ██║███████╗███████╗██████╔╝██║██║  ██║
		╚═╝     ╚═╝╚══════╝╚══════╝╚═════╝ ╚═╝╚═╝  ╚═╝
`
	author  = `		Auteur : `
	version = `		Version : `
)

func DrawStart() {
	defer time.Sleep(1 * time.Second)

	loger.Start(start)
	loger.Start(author+config.Author, "     ", version+config.Version)

	rgb.HiGreen.Println(start)
	rgb.HiGreen.Print(author+config.Author, "     ", version+config.Version)

}
