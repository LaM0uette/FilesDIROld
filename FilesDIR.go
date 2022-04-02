package main

import (
	"Test/build"
	"flag"
	"fmt"
	"os"
)

func main() {

	// setup flag for insert data of search in cli
	schMode := flag.String("m", "%", "Mode de recherche.")
	schWord := flag.String("f", "", "Non de fichier.")
	schExt := flag.String("e", "*", "Extension de fichier.")
	schPath := flag.String("l", build.CurrentDir(), "Extension de fichier.")
	flag.Parse()

	// generated the structure with data to search for files
	s := build.Search{
		Mode:      *schMode,
		Word:      *schWord,
		Extension: *schExt,
		Path:      *schPath,
	}

	// print on screen the start of program
	build.DrawStart()

	// print on screen the start of search
	err := s.SearchFiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
