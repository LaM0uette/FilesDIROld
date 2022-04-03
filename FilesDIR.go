package main

import (
	"Test/build"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	// setup flag for insert data of search in cli
	schCli := flag.Bool("r", false, "CLI / Run")
	schMode := flag.String("mode", "%", "Mode de recherche")
	schWord := flag.String("word", "", "Non de fichier")
	schExt := flag.String("ext", "*", "Extension de fichier")
	schMaj := flag.Bool("maj", false, "Autorise les majuscules")
	schPath := flag.String("path", build.CurrentDir(), "Chemin de recherche")
	flag.Parse()

	Save := build.DesktopDir()

	// print on screen the start of program
	build.DrawStart()

	// if is not mode cli, the user need to fill the settings of search
	if !*schCli {
		fmt.Print("Mode de recherche : ")
		*schMode, _ = reader.ReadString('\n')
		*schMode = strings.TrimSpace(*schMode)

		fmt.Print("Non de fichier : ")
		*schWord, _ = reader.ReadString('\n')
		*schWord = strings.TrimSpace(*schWord)

		fmt.Print("Extension de fichier : ")
		*schExt, _ = reader.ReadString('\n')
		*schExt = strings.TrimSpace(*schExt)

		fmt.Print("Chemin de recherche : ")
		*schPath, _ = reader.ReadString('\n')
		*schPath = strings.TrimSpace(*schPath)

		fmt.Print("\n\n")
		Save = build.CurrentDir()
	}

	// generated the structure with data to search for files
	s := build.Search{
		Mode:      *schMode,
		Word:      *schWord,
		Extension: *schExt,
		Maj:       *schMaj,
		Path:      *schPath,
		Save:      Save,
	}

	// print on screen the start of search
	err := s.SearchFiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *schCli {
		fmt.Print("Appuyer sur Entrée pour quitter...")
		_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err != nil {
			return
		}
	}

	os.Exit(1)
}
