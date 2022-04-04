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
	flagRunCLI := flag.Bool("r", false, "CLI / Run")
	flagMode := flag.String("mode", "%", "Mode de recherche")
	flagWord := flag.String("word", "", "Non de fichier")
	flagExt := flag.String("ext", "*", "Ext de fichier")
	flagMaj := flag.Bool("maj", false, "Autorise les majuscules")
	flagPath := flag.String("path", build.CurrentDir(), "Chemin de recherche")
	flag.Parse()

	// desktop dir is a default save folder
	saveFolder := build.DesktopDir()

	build.DrawStart()

	// if is not in cli mode, the user need to fill the settings of search
	if !*flagRunCLI {
		fmt.Print("Mode de recherche ( = | % | ^ | $ ) : ")
		*flagMode, _ = reader.ReadString('\n')
		*flagMode = strings.TrimSpace(*flagMode)

		fmt.Print("Non de fichier ( fiche, test, ...) : ")
		*flagWord, _ = reader.ReadString('\n')
		*flagWord = strings.TrimSpace(*flagWord)

		fmt.Print("Ext de fichier ( xlsx, jpg, ...) : ")
		*flagExt, _ = reader.ReadString('\n')
		*flagExt = strings.TrimSpace(*flagExt)

		fmt.Print("Chemin de recherche ( C:/... ) : ")
		*flagPath, _ = reader.ReadString('\n')
		*flagPath = strings.TrimSpace(*flagPath)

		fmt.Print("Prise en compte de la casse ( o | n ) : ")
		_schMaj, _ := reader.ReadString('\n')
		_schMaj = strings.TrimSpace(_schMaj)
		switch _schMaj {
		case "o":
			*flagMaj = true
		case "n":
			*flagMaj = false
		}

		fmt.Print("\n\n")
		saveFolder = build.CurrentDir()
	}

	// generated the structure with data to search for files
	s := build.Search{
		Mode: *flagMode,
		Word: *flagWord,
		Ext:  *flagExt,
		Maj:  *flagMaj,
		Path: *flagPath,
		Save: saveFolder,
	}

	// print on screen the start of search
	err := s.SearchFiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *flagRunCLI {
		fmt.Print("Appuyer sur Entrée pour quitter...")
		_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err != nil {
			return
		}
	}

	os.Exit(1)
}
