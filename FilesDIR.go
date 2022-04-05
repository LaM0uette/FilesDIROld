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
var saveFolder = build.DesktopDir()

func main() {

	build.DrawStart()

	// setup flag for insert data of search in cli
	flagReq := flag.String("req", "", "CLI / Run")

	flagRunCLI := flag.Bool("r", false, "CLI / Run")
	flagMode := flag.String("mode", "%", "Mode de recherche")
	flagWord := flag.String("word", "", "Non de fichier")
	flagExt := flag.String("ext", "*", "Ext de fichier")
	flagMaj := flag.Bool("maj", false, "Autorise les majuscules")
	flagSave := flag.Bool("save", false, "Sauvegarde à chaque fichier")
	flagPath := flag.String("path", build.CurrentDir(), "Chemin de recherche")
	flag.Parse()

	if *flagReq != "" {
		mode, word, ext, maj, save, err := build.ReadExcelFileForReq(*flagReq)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		*flagRunCLI = true
		*flagMode = mode
		*flagWord = word
		*flagExt = ext
		*flagMaj = maj
		*flagSave = save
		*flagPath = build.CurrentDir()
	}

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
		Mode:       *flagMode,
		Word:       *flagWord,
		Ext:        *flagExt,
		Maj:        *flagMaj,
		Save:       *flagSave,
		Path:       *flagPath,
		SaveFolder: saveFolder,
	}

	build.DrawStartSearch()

	reqUse, savePath, nbFolderMade, id, err := s.SearchFiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	build.DrawEndSearch(s.Path, reqUse, savePath, nbFolderMade, id)

	if *flagRunCLI {
		fmt.Print("Appuyer sur Entrée pour quitter...")
		_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err != nil {
			return
		}
	}
	os.Exit(1)
}
