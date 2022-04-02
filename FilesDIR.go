package main

import (
	"Test/build"
	"flag"
	"fmt"
	"os"
)

func main() {

	schMode := flag.String("m", "%", "Mode de recherche.")
	schWord := flag.String("f", "", "Non de fichier.")
	schExt := flag.String("e", "*", "Extension de fichier.")
	schPath := flag.String("l", build.CurrentDir(), "Extension de fichier.")
	flag.Parse()

	s := build.Search{
		Mode:      *schMode,
		Word:      *schWord,
		Extension: *schExt,
		Path:      *schPath,
	}

	build.DrawStart()

	err := s.SearchFiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
